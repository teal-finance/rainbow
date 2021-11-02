// Teal.Finance/Server is an opinionated boilerplate API and website server.
// Copyright (C) 2021 Teal.Finance contributors
//
// This file is part of Teal.Finance/Server, licensed under LGPL-3.0-or-later.
//
// Teal.Finance/Server is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License
// either version 3 of the License, or (at your option) any later version.
//
// Teal.Finance/Server is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.

package limiter

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/teal-finance/server/reserr"
	"golang.org/x/time/rate"
)

type ReqLimiter struct {
	visitors       map[string]*visitor
	defaultLimiter *rate.Limiter
	mu             sync.Mutex
	resErr         reserr.ResErr
}

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func New(maxReqBurst, maxReqPerMinute int, devMode bool, resErr reserr.ResErr) ReqLimiter {
	if devMode {
		maxReqBurst *= 10
		maxReqPerMinute *= 10
	}

	ratePerSecond := float64(maxReqPerMinute) / 60

	return ReqLimiter{
		visitors:       make(map[string]*visitor),
		defaultLimiter: rate.NewLimiter(rate.Limit(ratePerSecond), maxReqBurst),
		mu:             sync.Mutex{},
		resErr:         resErr,
	}
}

func (rl *ReqLimiter) Limit(next http.Handler) http.Handler {
	log.Printf("Middleware RateLimiter: burst=%v rate=%v/s",
		rl.defaultLimiter.Burst(), rl.defaultLimiter.Limit())

	go rl.removeOldVisitors()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("WARN: SplitHostPort(%v) %v", r.RemoteAddr, err)
			rl.resErr.Write(w, r, http.StatusInternalServerError, "Internal Server Error #3")

			return
		}

		limiter := rl.getVisitor(ip)
		if !limiter.Allow() {
			log.Print("TooManyRequests ", ip)
			rl.resErr.Write(w, r, http.StatusTooManyRequests, "Too Many Requests")

			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *ReqLimiter) removeOldVisitors() {
	for ; true; <-time.NewTicker(1 * time.Minute).C {
		rl.mu.Lock()

		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(rl.visitors, ip)
			}
		}

		rl.mu.Unlock()
	}
}

func (rl *ReqLimiter) getVisitor(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		rl.visitors[ip] = &visitor{rl.defaultLimiter, time.Now()}

		return rl.defaultLimiter
	}

	v.lastSeen = time.Now()

	return v.limiter
}
