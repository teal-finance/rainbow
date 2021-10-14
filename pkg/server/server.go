// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/open-policy-agent/opa/ast"
)

type Server struct {
	Version string
	DocURL  string // https://rainbow.teal.finance/doc

	DevMode bool

	HTTPPort int // port of the main server (REST API or any HTTP web server)
	ExpPort  int // export port for Prometheus

	MaxReqBurst     int
	MaxReqPerMinute int
	ratePerSecond   float64

	mu       sync.Mutex
	visitors map[string]*visitor

	httpConn     int64 // gauge
	httpActive   int64 // counter
	httpIdle     int64 // counter
	httpHijacked int64 // counter

	OPAFilenames []string
	opaCompiler  *ast.Compiler
}

// run the server for API endpoints.
func (s *Server) RunServer(h http.Handler) error {
	middlewares, connState := s.metricsServer()

	middlewares.Append(logRequests, s.limitReqRate, s.setServerHeader)

	if len(s.OPAFilenames) > 0 {
		err := s.loadPolicies()
		if err != nil {
			return err
		}

		middlewares.Append(s.auth)
	}

	middlewares.Append(allowCORS)

	port := strconv.Itoa(s.HTTPPort)

	log.Print("HTTP server listening on http://localhost:", port)

	server := http.Server{
		Addr:              ":" + port,
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

	go s.removeOldVisitors()

	return nil
}

func (s *Server) ReqError(w http.ResponseWriter, r *http.Request, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var err error

	switch {
	case r == nil:
		_, err = fmt.Fprintf(w, `{"error":%q, "doc":"`+s.DocURL+`"}`, errMsg)

	case r.URL.RawQuery == "":
		_, err = fmt.Fprintf(w, `{"error":%q, "path":%q, "doc":"`+s.DocURL+`"}`, errMsg, r.URL.Path)

	default:
		_, err = fmt.Fprintf(w, `{"error":%q, "path":%q, "query":%q, "doc":"`+s.DocURL+`"}`, errMsg, r.URL.Path, r.URL.RawQuery)
	}

	if err != nil {
		log.Printf("Write url=%v err: %v", r.URL.Path, err)
	}
}

func allowCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, req)
	})
}
