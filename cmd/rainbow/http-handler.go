// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/teal-finance/rainbow/pkg/all"
	"github.com/teal-finance/rainbow/pkg/server"
)

var optionsJSON []byte

func alwaysCollectOptions() {
	optionsJSON = []byte(`{"error":"initializing"}`)

	for {
		options, err := all.OptionsFromAllProviders()
		if err == nil {
			if b, err := json.Marshal(options); err != nil {
				log.Print("ERROR JSON Encode: ", err)
			} else {
				optionsJSON = b
			}
		}

		time.Sleep(10 * time.Minute)
	}
}

func replyOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(optionsJSON)
}

// apiHandler creates the mapping between the endpoints and the handler functions.
func apiHandler(s *server.Server) http.Handler {
	r := chi.NewRouter()

	// API
	r.Mount("/v0", apiRouter(s))

	return r
}

// apiRouter handles API endpoints.
func apiRouter(s *server.Server) chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", reserved(s))
		r.Get("/options", replyOptions)
	})

	return r
}

func reserved(s *server.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		s.ReqError(w, r, "Path is not valid. Please refer to the API doc.")
	}
}
