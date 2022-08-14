// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/teal-finance/garcon"
	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/api"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

func main() {
	parseFlags()

	g := garcon.New(
		garcon.WithURLs(*mainAddr),
		garcon.WithDev(*dev))

	providers := provider.AllProviders(*alert, g.ServerName.String())
	service := rainbow.NewService(providers, dbram.NewDB())

	// start the service in background
	go service.Run()

	server := server(&service, g)

	// delete secrets (no longer required)
	aes = nil
	hmac = nil

	log.Print("Server listening on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func server(service *rainbow.Service, g *garcon.Garcon) http.Server {
	chain, connState := g.StartMetricsServer(*expPort)
	chain = chain.Append(g.MiddlewareRejectUnprintableURI())
	chain = chain.Append(g.MiddlewareLogRequest())
	chain = chain.Append(g.MiddlewareRateLimiter(*reqBurst, *reqPerMinute))
	chain = chain.Append(g.MiddlewareServerHeader("Rainbow"))
	chain = chain.Append(g.MiddlewareCORS())

	r := handler(service, g)
	h := chain.Then(r)

	return http.Server{
		Addr:              listenAddr,
		Handler:           h,
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Minute, // Garcon.MiddlewareRateLimiter() delays responses, so people (attackers) who click frequently will wait longer.
		IdleTimeout:       time.Second,
		ConnState:         connState,
		ErrorLog:          log.Default(),
	}
}

func handler(s *rainbow.Service, g *garcon.Garcon) http.Handler {
	var ck garcon.TokenChecker
	if len(*aes) > 0 {
		ck = g.IncorruptibleChecker(*aes, 0, true)
	} else {
		ck = g.JWTChecker(*hmac, "FreePlan", 10, "PremiumPlan", 100)
	}

	r := chi.NewRouter()

	// Static website: set the cookie only when visiting index.html
	ws := g.NewStaticWebServer(*wwwDir)
	r.Get("/favicon.ico", ws.ServeFile("favicon.ico", "image/x-icon"))
	r.Get("/favicon.png", ws.ServeFile("favicon.png", "image/png"))
	r.Get("/preview.jpg", ws.ServeFile("preview.jpg", "image/jpeg"))
	r.With(ck.Chk).Get("/js/*", ws.ServeDir("text/javascript; charset=utf-8"))
	r.With(ck.Chk).Get("/assets/*", ws.ServeAssets())
	r.With(ck.Chk).Get("/version", garcon.ServeVersion())
	r.With(ck.Chk).Get("/version/", garcon.ServeVersion())

	// NotFound catches index.html and other Vue sub-folders
	r.With(ck.Set).NotFound(ws.ServeFile("index.html", "text/html; charset=utf-8"))

	// Disable the contact-form endpoint until it is well protected (CSRF).
	if false {
		// Forward submitted contact-form to Mattermost, and redirect browser to "/about".
		cf := g.NewContactForm("/about")
		r.With(ck.Chk).Post("/submit", cf.Notify(*form))
	}

	// API routes
	r.Route("/v0", func(r chi.Router) {
		h := api.NewHandler(s)

		// HTTP API
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

		r.With(ck.Chk).Get("/bff/cp", h.CallPut)

		// GraphQL API (and interactive API in developer mode)
		r.With(ck.Chk).Mount("/graphql", h.GraphQLHandler())
		if *dev {
			r.Mount("/graphiql", api.InteractiveGQLHandler("/v0/graphql"))
		}
	})

	return r
}
