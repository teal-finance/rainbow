// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/teal-finance/rainbow/pkg/all"
	"github.com/teal-finance/rainbow/pkg/server"
	"github.com/teal-finance/rainbow/pkg/server/fileserver"
)

const (
	queryDuration = 40 * time.Second
	sleepDuration = 10 * time.Minute
)

var (
	optionsJSON  = []byte(`{"error":"initializing"}`)
	lastModified string
	expires      string
)

func collectOptionsIndefinitely() {
	loopDuration := queryDuration + sleepDuration
	now := time.Now().UTC()

	for {
		beginTime := now

		options, err := all.OptionsFromAllProviders()
		if err == nil {
			if b, err := json.Marshal(options); err != nil {
				log.Print("ERROR JSON Encode: ", err)
			} else if !bytes.Equal(optionsJSON, b) {
				optionsJSON = b
				lastModified = beginTime.Format(http.TimeFormat)
			}
		}

		// try to estimate the next refresh time
		expires = time.Now().UTC().Add(loopDuration).Format(http.TimeFormat)

		time.Sleep(sleepDuration)

		now = time.Now().UTC()
		loopDuration += now.Sub(beginTime)
		loopDuration /= 2 // compute a basic average
	}
}

func replyOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if lastModified != "" {
		w.Header().Set("Last-Modified", lastModified)
		w.Header().Set("Expires", expires)
	}

	_, _ = w.Write(optionsJSON)
}

// apiHandler creates the mapping between the endpoints and the handler functions.
func apiHandler(s *server.Server) http.Handler {
	r := chi.NewRouter()

	// API
	r.Mount("/v0", apiRouter(s))

	// Static website files
	fs := fileserver.New(*wwwDir)
	r.NotFound(fs.ServeFile("index.html", "text/html; charset=utf-8")) // catch index.html and /Spots/BTC
	r.Get("/favicon.png", fs.ServeFile("favicon.png", "image/x-icon"))
	r.Get("/assets/js/*", fs.ServeDir("text/javascript; charset=utf-8"))
	r.Get("/assets/css/*", fs.ServeDir("text/css; charset=utf-8"))
	r.Get("/font/*", fs.ServeDir("font/woff2"))
	r.Get("/images/*", fs.ServeImages())

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
		s.Resp.Error(w, r, "Path is not valid. Please refer to the API doc.")
	}
}
