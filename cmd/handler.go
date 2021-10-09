// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// apiHandler creates the mapping between the endpoints and the handler functions.
func apiHandler() http.Handler {
	r := chi.NewRouter()

	// API
	r.Mount("/v1", apiRouter())

	return r
}

// apiRouter handles API endpoints.
func apiRouter() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", reserved)
		r.Get("/options", replyOptions)
	})

	return r
}

func reserved(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	reqError(w, r, "Path is not valid. Please refer to the API doc.")
}
