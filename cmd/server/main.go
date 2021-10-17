// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/teal-finance/rainbow/pkg/handler"
	"github.com/teal-finance/rainbow/pkg/server"
	"github.com/teal-finance/rainbow/pkg/server/cors"
	"github.com/teal-finance/rainbow/pkg/server/export"
	"github.com/teal-finance/rainbow/pkg/server/limiter"
	"github.com/teal-finance/rainbow/pkg/server/opa"
	"github.com/teal-finance/rainbow/pkg/server/resperr"
	"github.com/teal-finance/rainbow/pkg/service"
)

const version = "Rainbow-0.2.0"

func main() {
	parseFlags()

	// Start background service
	service := &service.Service{}
	go service.CollectOptionsIndefinitely()

	// Uniformize error responses with API doc
	respError := resperr.New("https://rainbow.teal.finance/doc")

	// Optionally it also starts a metrics server in background (if export port > 0).
	// The metrics server is for use with Prometheus or another compatible monitoring tool.
	metrics := export.Metrics{}
	middlewares, connState := metrics.StartServer(*expPort, *dev)

	// Limit the input request rate per IP
	reqLimiter := limiter.New(*maxReqBurst, *maxReqPerMinute, *dev, respError)
	middlewares.Append(server.LogRequests, reqLimiter.Limit, server.Header(version))

	// Set authentication policies
	if len(opaFilenames) > 0 {
		compiler, err := opa.LoadPolicies(opaFilenames)
		if err != nil {
			log.Fatal(err)
		}

		policy := opa.Policy{Compiler: compiler, Resp: respError}
		middlewares.Append(policy.Auth)
	}

	// CORS
	allowedOrigins := []string{"http://teal.finance:33322"}
	if *dev {
		allowedOrigins = append(allowedOrigins, "http://localhost:")
	}

	middlewares.Append(cors.HandleCORS(allowedOrigins))

	// h handles both the REST API and the static web files
	h := handler.Handler(service, respError, *wwwDir)
	h = middlewares.Then(h)

	runMainServer(h, connState)
}

// runMainServer runs in foreground the main server.
func runMainServer(h http.Handler, connState func(net.Conn, http.ConnState)) {
	addr := ":" + strconv.Itoa(*mainPort)

	server := http.Server{
		Addr:              addr,
		Handler:           h,
		TLSConfig:         nil,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
		MaxHeaderBytes:    222,
		TLSNextProto:      nil,
		ConnState:         connState,
		ErrorLog:          log.Default(),
		BaseContext:       nil,
		ConnContext:       nil,
	}

	log.Print("HTTP server listening on http://localhost", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Print("ERROR: Install ncat and ss: sudo apt install ncat iproute2")
		log.Printf("ERROR: Try to listen port %v: sudo ncat -l %v", *mainPort, *mainPort)
		log.Printf("ERROR: Get the process using port %v: sudo ss -pan | grep %v", *mainPort, *mainPort)
		log.Fatal(err)
	}
}
