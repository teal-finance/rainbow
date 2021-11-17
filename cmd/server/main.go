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

	g, err := garcon.New(
		garcon.WithOrigins(*mainAddr),
		garcon.WithDocURL("https:teal.finance/rainbow/doc"),
		garcon.WithServerHeader("Rainbow-v0.3"),
		garcon.WithLimiter(*reqBurst, *reqPerMinute),
		garcon.WithProm(*expPort),
		garcon.WithDev(*dev),
	)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr:              listenAddr,
		Handler:           g.Middlewares.Then(handler(&service, g)),
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
		ConnState:         g.ConnState,
		ErrorLog:          log.Default(),
	}

	log.Print("Server listening on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func handler(s *rainbow.Service, g *garcon.Garcon) http.Handler {
	r := chi.NewRouter()

	r.With(g.JWTChecker.SetCookie).Mount("/", webserver.Handler(g, *wwwDir))
	r.With(g.JWTChecker.ChkCookie).Mount("/v0", s.Handler())

	return r
}
