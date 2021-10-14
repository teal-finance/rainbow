// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package server

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/armon/go-metrics"
	"github.com/armon/go-metrics/prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// metricsServer creates and starts the Prometheus export server.
func (s *Server) metricsServer() (middlewares Chain, connState func(net.Conn, http.ConnState)) {
	if s.ExpPort <= 0 {
		log.Print("Disable Prometheus, export port=", s.ExpPort)
		return middlewares, nil
	}

	addr := ":" + strconv.Itoa(s.ExpPort)
	handler := metricsHandler()

	go func() {
		err := http.ListenAndServe(addr, handler)
		log.Fatal(err)
	}()

	log.Print("Prometheus export http://localhost" + addr)

	// connState counts the HTTP client connections.
	// In prod mode, we do not care about minor counting errors, we use the unsafe-thread version.
	// In dev mode we use the atomic version to avoid warnings from "go build -race".
	if s.DevMode {
		connState = s.updateConnCountersAtomic()
	} else {
		connState = s.updateConnCounters()
	}

	return NewChain(countRED), connState
}

// metricsHandler returns the endpoints "/metrics/xxx".
func metricsHandler() http.Handler {
	sink, err := prometheus.NewPrometheusSink()
	if err != nil {
		log.Print("ERROR: NewPrometheusSink cannot register sink because ", err)
		return nil
	}

	if _, err := metrics.NewGlobal(metrics.DefaultConfig("Rainbow"), sink); err != nil {
		log.Print("ERROR: Prometheus export is not able to provide metrics because ", err)
		return nil
	}

	handler := chi.NewRouter()
	handler.Handle("/metrics", promhttp.Handler())

	return handler
}

// countRED increments/decrements the RED metrics depending on incoming requests and outgoing responses.
func countRED(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		record := &statusRecorder{ResponseWriter: w, Status: "success"}

		next.ServeHTTP(record, r)

		labels := []metrics.Label{
			{Name: "method", Value: r.Method},
			{Name: "route", Value: r.RequestURI},
			{Name: "status", Value: record.Status},
		}

		duration := time.Since(start)
		metrics.AddSampleWithLabels([]string{"request_duration"}, float32(duration.Milliseconds()), labels)

		log.Printf("out %v %v %v", r.Method, r.URL, duration)
	})
}

// updateConnCounters increments/decrements the number of connections.
func (s *Server) updateConnCounters() func(net.Conn, http.ConnState) {
	return func(nc net.Conn, cs http.ConnState) {
		switch cs {
		// StateNew: the client just connects to TealAPI expecting a request.
		// Transition to either StateActive or StateClosed.
		case http.StateNew:
			s.httpConn++
		// StateActive when 1 or more bytes of a request has been read.
		// After the request is handled, transitions to StateClosed, StateHijacked, or StateIdle.
		// HTTP/2: StateActive only transitions away once all active requests are complete.
		case http.StateActive:
			s.httpActive++
		// StateIdle when handling a request is finished and is in the keep-alive state,
		// waiting for a new request, then transitions to either StateActive or StateClosed.
		case http.StateIdle:
			s.httpIdle++
		// StateHijacked is a terminal state: does not transition to StateClosed.
		case http.StateHijacked:
			s.httpHijacked++
			s.httpConn--
		// StateClosed is a terminal state.
		case http.StateClosed:
		default:
			s.httpConn--
		}
	}
}

func (s *Server) updateConnCountersAtomic() func(net.Conn, http.ConnState) {
	return func(nc net.Conn, cs http.ConnState) {
		switch cs {
		case http.StateNew:
			atomic.AddInt64(&s.httpConn, 1)
		case http.StateActive:
			atomic.AddInt64(&s.httpActive, 1)
		case http.StateIdle:
			atomic.AddInt64(&s.httpIdle, 1)
		case http.StateHijacked:
			atomic.AddInt64(&s.httpHijacked, 1)
			atomic.AddInt64(&s.httpConn, -1)
		case http.StateClosed:
		default:
			atomic.AddInt64(&s.httpConn, -1)
		}
	}
}

type statusRecorder struct {
	http.ResponseWriter
	Status string
}

func (r *statusRecorder) WriteHeader(status int) {
	if status != http.StatusOK {
		r.Status = "error"
	}

	r.ResponseWriter.WriteHeader(status)
}

// setServerHeader sets the Server HTTP header in the response.
func (s *Server) setServerHeader(next http.Handler) http.Handler {
	log.Print("Middleware: Set response header: Server ", s.Version)

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Server", s.Version)
			next.ServeHTTP(w, r)
		})
}

// logRequests logs the incoming HTTP requests.
func logRequests(next http.Handler) http.Handler {
	log.Print("Middleware: Will log all incoming request URL and remote address")

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("in  %v %v %v", r.Method, r.RequestURI, r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
}
