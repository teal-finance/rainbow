// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/garcon/webserver"
)

// WebHandler serves the static web files of the front-end.
func WebHandler(g *garcon.Garcon, wwwDir string) http.Handler {
	r := chi.NewRouter()

	// Static website files
	web := webserver.WebServer{Dir: wwwDir, ResErr: g.ResErr}
	r.NotFound(web.ServeFile("index.html", "text/html; charset=utf-8")) // catch index.html and other Vue sub-folders
	r.Get("/favicon.ico", web.ServeFile("favicon.ico", "image/x-icon"))
	r.Get("/favicon.png", web.ServeFile("favicon.png", "image/png"))
	r.Get("/preview.jpg", web.ServeFile("preview.jpg", "image/jpeg"))
	r.Get("/js/*", web.ServeDir("text/javascript; charset=utf-8"))
	r.Get("/assets/*", web.ServeAssets())

	return r
}
