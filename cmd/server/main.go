// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lynxai-team/emo"
	"github.com/lynxai-team/garcon/gc"
	"github.com/lynxai-team/garcon/gg"
	"github.com/lynxai-team/garcon/vv"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/api"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

var log = emo.NewZone("main")

func main() {
	emo.GlobalTimestamp(true)
	parseFlags()

	names := listProviderNames()
	log.Init("Providers:", names)

	g := gc.New(
		gc.WithURLs(gg.SplitClean(*mainAddr)...),
		gc.WithDev(*dev))

	// start the service in background
	providers := provider.Select(names, g.ServerName.String(), *alert)
	service := rainbow.NewService(providers, dbram.NewDB())
	go service.Run(*period)

	// chain middleware
	middleware, connState := g.StartExporter(*expPort)
	middleware = middleware.Append(
		g.MiddlewareRejectUnprintableURI(),
		g.MiddlewareLogRequest(),
		g.MiddlewareRateLimiter(*reqBurst, *reqPerMinute),
		g.MiddlewareCORS(),
		g.MiddlewareServerHeader("Rainbow"))

	// middleware to set/check cookies
	var ck gc.TokenChecker
	if len(*aes) > 0 {
		ck = g.IncorruptibleChecker(*aes, 0, true)
	} else {
		ck = g.JWTChecker(*hmac, "FreePlan", 10, "PremiumPlan", 100)
	}

	// delete secrets (no longer required)
	aes = nil
	hmac = nil

	router := chi.NewRouter()

	// Static website
	ws := g.NewStaticWebServer(*wwwDir)
	// set the cookie when visiting index.html (NotFound = catches index.html and other Vue sub-folders)
	router.With(ck.Set).NotFound(ws.ServeFile("index.html", "text/html; charset=utf-8"))
	// protect web files: reject invalid cookie
	router.With(ck.Chk).Get("/js/*", ws.ServeDir("text/javascript; charset=utf-8"))
	router.With(ck.Chk).Get("/assets/*", ws.ServeAssets())
	router.With(ck.Chk).Get("/version", vv.ServeVersion())
	// do not protect favicon and other public images
	router.Get("/favicon.ico", ws.ServeFile("favicon.ico", "image/x-icon"))
	router.Get("/favicon.png", ws.ServeFile("favicon.png", "image/png"))
	router.Get("/img/*", ws.ServeImages())

	// Disable the contact-form endpoint until we protect it against DoS
	if false {
		contactForm := g.NewContactForm("/about")                      // submitted contact-form redirects to "/about"
		router.With(ck.Chk).Post("/submit", contactForm.Notify(*form)) // forward contact-form to Mattermost
	}

	// API routes
	router.Route("/v0", func(r chi.Router) {
		h := api.NewHandler(&service)

		// HTTP API for Jupyter notebooks
		r.With(ck.Vet).Route("/options", func(r chi.Router) {
			r.HandleFunc("/", h.Options)
			r.HandleFunc("/{asset}", h.Options)
			r.HandleFunc("/{asset}/", h.Options)
			r.HandleFunc("/{asset}/{expiry}", h.Options)
			r.HandleFunc("/{asset}/{expiry}/", h.Options)
			r.HandleFunc("/{asset}/{expiry}/{provider}", h.Options)
			r.HandleFunc("/{asset}/{expiry}/{provider}/", h.Options)
			r.HandleFunc("/{asset}/{expiry}/{provider}/{format}", h.Options)
			r.HandleFunc("/{asset}/{expiry}/{provider}/{format}/", h.Options)
		})

		// HTTP API using the Backend-For-Frontend pattern
		r.With(ck.Vet).Get("/bff/cp", h.CallPut)

		// GraphQL API to display the screener on a web page
		r.With(ck.Vet).Mount("/graphql", h.GraphQLHandler())
		if *dev { // interactive GraphQL API in dev. mode only
			r.Mount("/graphiql", api.InteractiveGQLHandler("/v0/graphql"))
		}
	})

	server := http.Server{
		Addr:              listenAddr,
		Handler:           middleware.Then(router),
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Minute, // Garcon.MiddlewareRateLimiter() delays responses, so people (attackers) who click frequently will wait longer.
		IdleTimeout:       time.Second,
		ConnState:         connState,
		ErrorLog:          log.Default(),
	}

	log.Print("Server listening on http://localhost" + server.Addr)
	log.Fatal(server.ListenAndServe())
}
