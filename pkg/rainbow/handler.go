package rainbow

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	s Service
}

func (s Service) Handler() http.Handler {
	h := handler{s: s}

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/options", h.getOptions)
	})

	return r
}

func (h handler) getOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Last-Modified", lastUpdate.Format(time.RFC3339))

	if err := json.NewEncoder(w).Encode(options); err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}
}
