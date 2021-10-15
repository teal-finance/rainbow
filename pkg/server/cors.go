// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/cors"
)

// handleCORS uses restrictive CORS values.
func handleCORS(origins []string) func(next http.Handler) http.Handler {
	options := cors.Options{
		AllowedOrigins:     []string{},               // No need because use function
		AllowOriginFunc:    nil,                      // Function is set below
		AllowedMethods:     []string{http.MethodGet}, // The most restrictive
		AllowedHeaders:     []string{"Origin", "Accept", "Content-Type", "Authorization", "Cookie"},
		ExposedHeaders:     []string{},
		AllowCredentials:   true,
		MaxAge:             60,    // Cache the response for 1 minute https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age
		OptionsPassthrough: false, // false = this middleware stops OPTION requests
		Debug:              true,
	}

	for i, dns := range origins {
		origins[i] = insertSchema(dns)
	}

	if len(origins) == 1 {
		options.AllowOriginFunc = oneOrigin(origins[0])
	} else {
		options.AllowOriginFunc = multipleOrigins(origins)
	}

	log.Print("Middleware CORS: ", options)

	return cors.Handler(options)
}

func insertSchema(domain string) string {
	if !strings.HasPrefix(domain, "https://") &&
		!strings.HasPrefix(domain, "http://") {
		return "http://" + domain
	}

	return domain
}

func oneOrigin(addr string) func(r *http.Request, origin string) bool {
	log.Print("CORS: Set origin: ", addr)

	return func(r *http.Request, origin string) bool {
		return origin == addr
	}
}

func multipleOrigins(prefixes []string) func(r *http.Request, origin string) bool {
	log.Print("CORS: Set origin prefixes: ", prefixes)

	return func(r *http.Request, origin string) bool {
		for _, prefix := range prefixes {
			if strings.HasPrefix(origin, prefix) {
				log.Printf("CORS: Accept %v because starts with prefix %v", origin, prefix)
				return true
			}

			log.Printf("CORS: %v does not begin with %v", origin, prefix)
		}

		log.Printf("CORS: Refuse %v because different from %v", origin, prefixes)

		return false
	}
}
