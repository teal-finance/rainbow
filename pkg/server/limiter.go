// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package server

import (
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func (s *Server) removeOldVisitors() {
	for ; true; <-time.NewTicker(1 * time.Minute).C {
		s.mu.Lock()

		for ip, v := range s.visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(s.visitors, ip)
			}
		}

		s.mu.Unlock()
	}
}

func (s *Server) getVisitor(ip string) *rate.Limiter {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, exists := s.visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(rate.Limit(s.ratePerSecond), s.MaxReqBurst)
		s.visitors[ip] = &visitor{limiter, time.Now().UTC()}

		return limiter
	}

	v.lastSeen = time.Now().UTC()

	return v.limiter
}

func (s *Server) limitReqRate(next http.Handler) http.Handler {
	s.visitors = make(map[string]*visitor)
	s.ratePerSecond = float64(s.MaxReqPerMinute) / 60
	log.Print("Enable Rate limiter rate=", s.ratePerSecond)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("WARN: SplitHostPort(%v) %v", r.RemoteAddr, err)
			w.WriteHeader(http.StatusInternalServerError)
			s.ReqError(w, r, "Internal Server Error")
			return
		}

		limiter := s.getVisitor(ip)
		if !limiter.Allow() {
			log.Print("TooManyRequests IP ", ip)
			w.WriteHeader(http.StatusTooManyRequests)
			s.ReqError(w, r, "Too Many Requests")
			return
		}

		next.ServeHTTP(w, r)
	})
}
