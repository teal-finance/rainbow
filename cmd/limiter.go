// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
	mu       sync.Mutex
	visitors = make(map[string]*visitor)
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func removeOldVisitors() {
	for ; true; <-time.NewTicker(1 * time.Minute).C {
		mu.Lock()

		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}

		mu.Unlock()
	}
}

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		ratePerSecond := float64(*maxReqPerMinute) / 60
		limiter := rate.NewLimiter(rate.Limit(ratePerSecond), *maxReqBurst)
		visitors[ip] = &visitor{limiter, time.Now().UTC()}

		return limiter
	}

	v.lastSeen = time.Now().UTC()

	return v.limiter
}

func limitReqRate(next http.Handler) http.Handler {
	log.Print("Enable Rate limiter")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("WARN: SplitHostPort(%v) %v", r.RemoteAddr, err)
			w.WriteHeader(http.StatusInternalServerError)
			reqError(w, r, "Internal Server Error")
			return
		}

		limiter := getVisitor(ip)
		if !limiter.Allow() {
			log.Print("TooManyRequests IP ", ip)
			w.WriteHeader(http.StatusTooManyRequests)
			reqError(w, r, "Too Many Requests")
			return
		}

		next.ServeHTTP(w, r)
	})
}
