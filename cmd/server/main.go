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

	var tokenOption garcon.Option
	if len(*aes) > 0 {
		tokenOption = garcon.WithIncorruptible(*aes, 0, true)
	} else {
		tokenOption = garcon.WithJWT(*hmac, "FreePlan", 10, "PremiumPlan", 100)
	}
	// secrets no longer required => erase them
	aes = nil
	hmac = nil

	g, err := garcon.New(
		garcon.WithURLs(*mainAddr),
		garcon.WithDocURL("/doc"),
		garcon.WithServerHeader("Rainbow"),
		tokenOption,
		garcon.WithLimiter(*reqBurst, *reqPerMinute),
		garcon.WithProm(*expPort, *mainAddr),
		garcon.WithDev(*dev))
	if err != nil {
		log.Fatal(err)
	}

	providers := provider.AllProviders(*alert, g.Namespace.String())

	// Start the service in background
	service := rainbow.NewService(providers, dbram.NewDB())
	go service.Run()

	server := http.Server{
		Addr:              listenAddr,
		Handler:           g.Middlewares.Then(handler(&service, g)),
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Minute, // Garcon.Limiter postpones response, attacker should wait long time.
		IdleTimeout:       time.Second,
		ConnState:         g.ConnState,
		ErrorLog:          log.Default(),
	}

	log.Print("Server listening on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func handler(s *rainbow.Service, g *garcon.Garcon) http.Handler {
	r := chi.NewRouter()
	c := g.Checker

	// Static website: set the cookie only when visiting index.html
	ws := garcon.NewStaticWebServer(*wwwDir, g.ErrWriter)
	r.Get("/favicon.ico", ws.ServeFile("favicon.ico", "image/x-icon"))
	r.Get("/favicon.png", ws.ServeFile("favicon.png", "image/png"))
	r.Get("/preview.jpg", ws.ServeFile("preview.jpg", "image/jpeg"))
	r.With(c.Chk).Get("/js/*", ws.ServeDir("text/javascript; charset=utf-8"))
	r.With(c.Chk).Get("/assets/*", ws.ServeAssets())
	r.With(c.Chk).Get("/version", g.ServeVersion())
	r.With(c.Chk).Get("/version/", g.ServeVersion())

	// NotFound catches index.html and other Vue sub-folders
	r.With(c.Set).NotFound(ws.ServeFile("index.html", "text/html; charset=utf-8"))

	// Disable the contact-form endpoint until it is well protected (CSRF).
	if false {
		// Forward submitted contact-form to Mattermost, and redirect browser to "/about".
		cf := garcon.NewContactForm("/about", *form, g.ErrWriter)
		r.With(c.Chk).Post("/submit", cf.NotifyWebForm())
	}

	// API routes
	r.Route("/v0", func(r chi.Router) {
		h := api.NewHandler(s)

		// HTTP API
		r.With(g.Checker.Vet).Route("/options", func(r chi.Router) {
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

		r.With(g.Checker.Chk).Get("/bff/cp", h.CallPut)

		// GraphQL API (and interactive API in developer mode)
		r.With(g.Checker.Chk).Mount("/graphql", h.GraphQLHandler())
		if *dev {
			r.Mount("/graphiql", api.InteractiveGQLHandler("/v0/graphql"))
		}
	})

	return r
}
