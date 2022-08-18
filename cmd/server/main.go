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
	"github.com/spf13/viper"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/rainbow/pkg/provider/deltaexchange"
	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/provider/lyra"
	"github.com/teal-finance/rainbow/pkg/provider/psyoptions"
	"github.com/teal-finance/rainbow/pkg/provider/thales"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/api"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

func main() {
	parseFlags()

	// parse configuration file
	conf := Config{}
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("reading configuration file : %v", err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("decoding configuration : %v", err)
	}

	// garcon config
	g := garcon.New(
		garcon.WithURLs(*mainAddr),
		garcon.WithDev(*dev),
	)

	// start the service in background
	providers := []rainbow.Provider{
		psyoptions.Provider{},
		deribit.Provider{},
		lyra.Provider{},
		zerox.Provider{},
		zetamarkets.Provider{Endpoint: conf.RPC.Serum},
		thales.Provider{},
		deltaexchange.Provider{},
	}
	service := rainbow.NewService(providers, dbram.NewDB())
	go service.Run()

	// chain middleware
	middleware, connState := g.StartMetricsServer(*expPort)
	middleware = middleware.Append(g.MiddlewareRejectUnprintableURI(),
		g.MiddlewareLogRequest(),
		g.MiddlewareRateLimiter(*reqBurst, *reqPerMinute),
		g.MiddlewareServerHeader("Rainbow"),
		g.MiddlewareCORS())

	// middleware to set/check cookies
	var ck garcon.TokenChecker
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
	router.With(ck.Chk).Get("/version", garcon.ServeVersion())
	// do not protect favicon and preview.jpg
	router.Get("/favicon.ico", ws.ServeFile("favicon.ico", "image/x-icon"))
	router.Get("/favicon.png", ws.ServeFile("favicon.png", "image/png"))
	router.Get("/preview.jpg", ws.ServeFile("preview.jpg", "image/jpeg"))

	// Disable the contact-form endpoint until we protect it against DoS
	if false {
		cf := g.NewContactForm("/about")                      // submitted contact-form redirects to "/about"
		router.With(ck.Chk).Post("/submit", cf.Notify(*form)) // forward contact-form to Mattermost
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

	log.Print("Server listening on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}
