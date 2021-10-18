// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teal-finance/server/fileserver"
	"github.com/teal-finance/server/reserr"
)

// Handler creates the mapping between the endpoints and the handler functions.
func (s *Service) Handler(resErr reserr.ResErr, wwwDir string) http.Handler {
	r := chi.NewRouter()

	// API
	r.Mount("/v0", s.apiRouter(resErr))

	// Static website files
	fs := fileserver.FileServer{Dir: wwwDir, ResErr: resErr}
	r.NotFound(fs.ServeFile("index.html", "text/html; charset=utf-8")) // catch index.html and other Vue sub-folders
	r.Get("/favicon.png", fs.ServeFile("favicon.png", "image/x-icon"))
	r.Get("/assets/js/*", fs.ServeDir("text/javascript; charset=utf-8"))
	r.Get("/assets/css/*", fs.ServeDir("text/css; charset=utf-8"))
	r.Get("/font/*", fs.ServeDir("font/woff2"))
	r.Get("/images/*", fs.ServeImages())

	return r
}

// apiRouter handles API endpoints.
func (s *Service) apiRouter(resp reserr.ResErr) chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", resp.NotImplemented)
		r.Get("/options", s.ReplyOptions)
	})

	r.NotFound(resp.InvalidPath)

	return r
}
