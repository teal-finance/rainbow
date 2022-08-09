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

	var tokenCheckerOption garcon.Option
	if len(*aes) > 0 {
		tokenCheckerOption = garcon.WithIncorruptible(*aes, 0, true)
	} else {
		tokenCheckerOption = garcon.WithJWT(*hmac, "FreePlan", 10, "PremiumPlan", 100)
	}

	g := garcon.New(
		garcon.WithURLs(*mainAddr),
		garcon.WithDocURL("/doc"),
		garcon.WithServerHeader("Rainbow"),
		tokenCheckerOption,
		garcon.WithLimiter(*reqBurst, *reqPerMinute),
		garcon.WithProm(*expPort, *mainAddr),
		garcon.WithDev(*dev))

	providers := provider.AllProviders(*alert, g.Namespace.String())
	service := rainbow.NewService(providers, dbram.NewDB())

	// start the service in background
	go service.Run()

	server := server(&service, g)

	// delete secrets (no longer required)
	g.EraseSecretKey()
	aes = nil
	hmac = nil

	log.Print("Server listening on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func server(service *rainbow.Service, g *garcon.Garcon) http.Server {
	middleware, connState := g.ChainMiddleware()
	h := handler(service, g)
	h = middleware.Then(h)

	return http.Server{
		Addr:              listenAddr,
		Handler:           h,
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Minute, // Garcon.RateLimiter() delays responses, so people (attackers) who click frequently will wait longer.
		IdleTimeout:       time.Second,
		ConnState:         connState,
		ErrorLog:          log.Default(),
	}
}

func handler(s *rainbow.Service, g *garcon.Garcon) http.Handler {
	r := chi.NewRouter()
	c := g.TokenChecker()

	// Static website: set the cookie only when visiting index.html
	ws := garcon.NewStaticWebServer(*wwwDir, g.Writer)
	r.Get("/favicon.ico", ws.ServeFile("favicon.ico", "image/x-icon"))
	r.Get("/favicon.png", ws.ServeFile("favicon.png", "image/png"))
	r.Get("/preview.jpg", ws.ServeFile("preview.jpg", "image/jpeg"))
	r.With(c.Chk).Get("/js/*", ws.ServeDir("text/javascript; charset=utf-8"))
	r.With(c.Chk).Get("/assets/*", ws.ServeAssets())
	r.With(c.Chk).Get("/version", garcon.ServeVersion())
	r.With(c.Chk).Get("/version/", garcon.ServeVersion())

	// NotFound catches index.html and other Vue sub-folders
	r.With(c.Set).NotFound(ws.ServeFile("index.html", "text/html; charset=utf-8"))

	// Disable the contact-form endpoint until it is well protected (CSRF).
	if false {
		// Forward submitted contact-form to Mattermost, and redirect browser to "/about".
		cf := garcon.NewContactForm("/about", *form, g.Writer)
		r.With(c.Chk).Post("/submit", cf.NotifyWebForm())
	}

	// API routes
	r.Route("/v0", func(r chi.Router) {
		h := api.NewHandler(s)

		// HTTP API
		r.With(c.Vet).Route("/options", func(r chi.Router) {
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

		r.With(c.Chk).Get("/bff/cp", h.CallPut)

		// GraphQL API (and interactive API in developer mode)
		r.With(c.Chk).Mount("/graphql", h.GraphQLHandler())
		if *dev {
			r.Mount("/graphiql", api.InteractiveGQLHandler("/v0/graphql"))
		}
	})

	return r
}
