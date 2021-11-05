// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	s *Service
}

func (s *Service) Handler() http.Handler {
	h := handler{s: s}

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/options", h.getOptions)
		r.Get("/options/cp", h.getCPFormat)
	})

	return r
}

func (h handler) getOptions(w http.ResponseWriter, r *http.Request) {
	options, err := h.s.Options()
	if err != nil {
		log.Print("ERROR getOptions ", err)
		http.Error(w, "No Content", http.StatusNoContent)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(options); err != nil {
		log.Print("ERROR getOptions ", err)
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)

		return
	}
}

func (h handler) getCPFormat(w http.ResponseWriter, r *http.Request) {
	t, err := h.s.store.GetCPFormat()
	if err != nil {
		log.Print("ERROR getCPFormat ", err)
		http.Error(w, "No Content", http.StatusNoContent)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Print("ERROR getCPFormat ", err)
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)

		return
	}
}
