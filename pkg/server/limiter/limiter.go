// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package limiter

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/teal-finance/rainbow/pkg/server/resperr"
)

type ReqLimiter struct {
	visitors       map[string]*visitor
	defaultLimiter *rate.Limiter
	mu             sync.Mutex
	resp           resperr.RespErr
}

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func New(maxReqBurst, maxReqPerMinute int, devMode bool, resp resperr.RespErr) ReqLimiter {
	if devMode {
		maxReqBurst *= 10
		maxReqPerMinute *= 10
	}

	ratePerSecond := float64(maxReqPerMinute) / 60

	return ReqLimiter{
		visitors:       make(map[string]*visitor),
		defaultLimiter: rate.NewLimiter(rate.Limit(ratePerSecond), maxReqBurst),
		mu:             sync.Mutex{},
		resp:           resp,
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
			w.WriteHeader(http.StatusInternalServerError)
			rl.resp.Error(w, r, "Internal Server Error")
			return
		}

		limiter := rl.getVisitor(ip)
		if !limiter.Allow() {
			log.Print("TooManyRequests ", ip)
			w.WriteHeader(http.StatusTooManyRequests)
			rl.resp.Error(w, r, "Too Many Requests")
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
