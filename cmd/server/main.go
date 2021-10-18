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

	"github.com/teal-finance/server"
	"github.com/teal-finance/server/chain"
	"github.com/teal-finance/server/cors"
	"github.com/teal-finance/server/limiter"
	"github.com/teal-finance/server/metrics"
	"github.com/teal-finance/server/opa"
	"github.com/teal-finance/server/reserr"

	"github.com/teal-finance/rainbow/pkg/service"
)

const version = "Rainbow-0.2.0"

func main() {
	parseFlags()

	// Start the service in background
	service := &service.Service{}
	go service.CollectOptionsIndefinitely()

	// Uniformize error responses with API doc
	resErr := reserr.New("https://rainbow.teal.finance/doc")

	middlewares, connState := setMiddlewares(resErr)

	// h handles both the REST API and the static web files
	h := service.Handler(resErr, *wwwDir)
	h = middlewares.Then(h)

	runMainServer(h, connState)
}

func setMiddlewares(resErr reserr.ResErr) (middlewares chain.Chain, connState func(net.Conn, http.ConnState)) {
	// Start a metrics server in background if export port > 0.
	// The metrics server is for use with Prometheus or another compatible monitoring tool.
	metrics := metrics.Metrics{}
	middlewares, connState = metrics.StartServer(*expPort, *dev)

	// Limit the input request rate per IP
	reqLimiter := limiter.New(*maxReqBurst, *maxReqPerMinute, *dev, resErr)

	middlewares = middlewares.Append(
		server.LogRequests,
		server.Header(version),
		reqLimiter.Limit,
	)

	// Endpoint authentication rules (Open Policy Agent)
	// Set authentication policies
	if len(opaFilenames) > 0 {
		policy, err := opa.New(resErr, opaFilenames)
		if err != nil {
			log.Fatal(err)
		}

		middlewares = middlewares.Append(policy.Auth)
	}

	// CORS
	allowedOrigins := []string{"http://teal.finance:33322"}
	if *dev {
		allowedOrigins = append(allowedOrigins, "http://localhost:")
	}

	middlewares = middlewares.Append(cors.HandleCORS(allowedOrigins))

	return middlewares, connState
}

// runMainServer runs in foreground the main server.
func runMainServer(h http.Handler, connState func(net.Conn, http.ConnState)) {
	server := http.Server{
		Addr:              ":" + strconv.Itoa(*mainPort),
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

	log.Print("Server listening on http://localhost", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Print("ERROR: Install ncat and ss: sudo apt install ncat iproute2")
		log.Printf("ERROR: Try to listen port %v: sudo ncat -l %v", *mainPort, *mainPort)
		log.Printf("ERROR: Get the process using port %v: sudo ss -pan | grep %v", *mainPort, *mainPort)
		log.Fatal(err)
	}
}
