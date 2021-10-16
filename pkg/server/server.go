// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/open-policy-agent/opa/ast"

	"github.com/teal-finance/rainbow/pkg/server/cors"
	"github.com/teal-finance/rainbow/pkg/server/export"
	"github.com/teal-finance/rainbow/pkg/server/limiter"
	"github.com/teal-finance/rainbow/pkg/server/resperr"
)

type Server struct {
	Version string
	Resp    resperr.RespErr

	// CORS
	AllowedOrigins []string // used for CORS

	// OPA
	OPAFilenames []string
	compiler     *ast.Compiler

	metrics export.Metrics
}

// RunServer runs the HTTP server in foreground.
// Optionally it also starts a metrics server in background (if export port > 0).
// The metrics server is for use with Prometheus or another compatible monitoring tool.
func (s *Server) RunServer(h http.Handler, port, expPort, maxReqBurst, maxReqPerMinute int, devMode bool) error {
	middlewares, connState := s.metrics.StartServer(expPort, devMode)

	reqLimiter := limiter.New(maxReqBurst, maxReqPerMinute, devMode, s.Resp)

	middlewares.Append(logRequests, reqLimiter.Limit, s.setServerHeader)

	if len(s.OPAFilenames) > 0 {
		err := s.loadPolicies()
		if err != nil {
			return err
		}

		middlewares.Append(s.auth)
	}

	middlewares.Append(cors.HandleCORS(s.AllowedOrigins))

	addr := ":" + strconv.Itoa(port)

	log.Print("HTTP server listening on http://localhost", addr)

	// main server: REST API or any HTTP web server
	server := http.Server{
		Addr:              addr,
		Handler:           middlewares.Then(h),
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

	if err := server.ListenAndServe(); err != nil {
		log.Print("ERROR: Install ncat and ss: sudo apt install ncat iproute2")
		log.Printf("ERROR: Try to listen port %v: sudo ncat -l %v", port, port)
		log.Printf("ERROR: Get the process using port %v: sudo ss -pan | grep %v", port, port)

		return err
	}

	return nil
}

// setServerHeader sets the Server HTTP header in the response.
func (s *Server) setServerHeader(next http.Handler) http.Handler {
	log.Print("Middleware response HTTP header: Set Server ", s.Version)

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Server", s.Version)
			next.ServeHTTP(w, r)
		})
}

// logRequests logs the incoming HTTP requests.
func logRequests(next http.Handler) http.Handler {
	log.Print("Middleware logger: log requested URLs and remote addresses")

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("in  %v %v %v", r.Method, r.RequestURI, r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
}
