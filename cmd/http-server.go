// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func runAPIServer() {
	middlewares, connState := metricsServer()

	// run the server for API endpoints
	log.Print("API server listening on http://localhost:", *port)

	server := http.Server{
		Addr:              ":" + *port,
		Handler:           middlewares.Then(apiHandler()),
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
		log.Printf("ERROR: Try to listen port %v: sudo ncat -l %v", *port, *port)
		log.Printf("ERROR: Get the process using port %v: sudo ss -pan | grep %v", *port, *port)
		log.Fatal(err)
	}

	go removeOldVisitors()
}

func reqError(w http.ResponseWriter, r *http.Request, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var err error

	switch {
	case r == nil:
		_, err = fmt.Fprintf(w, `{"error":%q, "doc":"https://rainbow.TODO/doc"}`, errMsg)

	case r.URL.RawQuery == "":
		_, err = fmt.Fprintf(w, `{"error":%q, "path":%q, "doc":"https://rainbow.TODO/doc"}`, errMsg, r.URL.Path)

	default:
		_, err = fmt.Fprintf(w, `{"error":%q, "path":%q, "query":%q, "doc":"https://rainbow.TODO/doc"}`, errMsg, r.URL.Path, r.URL.RawQuery)
	}

	if err != nil {
		log.Printf("Write url=%v err: %v", r.URL.Path, err)
	}
}
