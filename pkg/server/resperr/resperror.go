// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package resperr

import (
	"fmt"
	"log"
	"net/http"
)

type RespErr string

// New is useless.
func New(docURL string) RespErr {
	return RespErr(docURL)
}

func (url RespErr) NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	url.Error(w, r, "Path is reserved for future use. Please contact us to share your ideas.")
}

func (url RespErr) InvalidPath(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	url.Error(w, r, "Path is not valid. Please refer to the API doc.")
}

func (url RespErr) Error(w http.ResponseWriter, r *http.Request, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var err error

	switch {
	case r == nil:
		_, err = fmt.Fprintf(w, `{"error":%q, "doc":%q}`, msg, url)

	case r.URL.RawQuery == "":
		_, err = fmt.Fprintf(w, `{"error":%q, "path":%q, "doc":%q}`, msg, r.URL.Path, url)

	default:
		_, err = fmt.Fprintf(w, `{"error":%q, "path":%q, "query":%q, "doc":%q}`, msg, r.URL.Path, r.URL.RawQuery, url)
	}

	if err != nil {
		log.Printf("Write url=%v err: %v", r.URL.Path, err)
	}
}

func Error(w http.ResponseWriter, r *http.Request, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var err error

	switch {
	case r == nil:
		_, err = fmt.Fprintf(w, `{"error":%q}`, msg)

	case r.URL.RawQuery == "":
		_, err = fmt.Fprintf(w, `{"error":%q, "path":%q}`, msg, r.URL.Path)

	default:
		_, err = fmt.Fprintf(w, `{"error":%q, "path":%q, "query":%q}`, msg, r.URL.Path, r.URL.RawQuery)
	}

	if err != nil {
		log.Printf("Write url=%v err: %v", r.URL.Path, err)
	}
}
