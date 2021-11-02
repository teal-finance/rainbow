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

// Handler to server static web frontend files
func Handler(resErr reserr.ResErr, wwwDir string) http.Handler {
	r := chi.NewRouter()

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
