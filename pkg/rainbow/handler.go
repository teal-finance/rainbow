// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"encoding/json"
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
		// Deprecated API
		r.Get("/options", h.getOptions)
		// New API
		r.Get("/expiries", h.getExpiries)
		r.Get("/tables/default", h.getTable)
		r.Get("/tables/{asset}/{expiry}", h.getTable)
	})

	return r
}

// deprecated API.
func (h handler) getOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	options, err := h.s.Options()
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(options); err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
}

// new API endpoints

func (h handler) getExpiries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	e, err := h.s.store.GetExpiries()
	if err != nil {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}

	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
}

func (h handler) getTable(w http.ResponseWriter, r *http.Request) {
	asset := chi.URLParam(r, "asset")
	expiry := chi.URLParam(r, "expiry")

	w.Header().Set("Content-Type", "application/json")

	t, err := h.s.store.GetTable(asset + expiry)
	if err != nil {
		http.Error(w, "No data for asset="+asset+" expiry="+expiry, http.StatusNoContent)
		return
	}

	if err := json.NewEncoder(w).Encode(t); err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
}
