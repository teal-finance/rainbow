// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/garcon/reserr"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
	"github.com/teal-finance/rainbow/pkg/webserver"
)

func main() {
	parseFlags()

	// Start the service in background
	service := rainbow.NewService(
		provider.AllProviders(),
		&dbram.DB{},
	)
	go service.Run()

	g := garcon.Garcon{
		Version:        "Rainbow-v0",
		ResErr:         "https:teal.finance/rainbow/doc",
		AllowedOrigins: []string{*mainAddr},
		OPAFilenames:   nil,
	}

	middlewares, connState, err := g.Setup(0, *expPort, *reqBurst, *reqPerMinute, *dev)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr:              listenAddr,
		Handler:           middlewares.Then(handler(&service)),
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
		ConnState:         connState,
		ErrorLog:          log.Default(),
	}

	log.Print("Server listening on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func handler(s *rainbow.Service) http.Handler {
	r := chi.NewRouter()

	r.Mount("/", webserver.Handler(reserr.New(*mainAddr+listenAddr+"/doc"), *wwwDir))
	r.Mount("/v0", s.Handler())

	return r
}
