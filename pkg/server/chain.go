// modified copy of https://github.com/justinas/alice
// MIT License
// Copyright (c) 2014 Justinas Stankevicius
// Copyright (c) 2021 teal.finance

package server

import (
	"net/http"
)

// Middleware is a constructor function returning a http.Handler.
type Middleware func(http.Handler) http.Handler

// Chain acts as a list of http.Handler middlewares.
// Chain is effectively immutable:
// once created, it will always hold
// the same set of middlewares in the same order.
type Chain []Middleware

// NewChain creates a new chain,
// memorizing the given list of middlewares.
// NewChain serves no other function,
// middlewares are only constructed upon a call to Then().
func NewChain(middlewares ...Middleware) Chain {
	return middlewares
}

// Append extends a chain, adding the specified middlewares
// as the last ones in the request flow.
//
//     chain := server.NewChain(m1, m2)
//     chain.Append(m3, m4)
//     // requests in chain go m1 -> m2 -> m3 -> m4
func (c *Chain) Append(middlewares ...Middleware) {
	*c = append(*c, middlewares...)
}

// Then chains the middleware and returns the final http.Handler.
//     server.NewChain(m1, m2, m3).Then(h)
// is equivalent to:
//     m1(m2(m3(h)))
//
// When the request comes in, it will be passed to m1, then m2, then m3
// and finally, the given handler
// (assuming every middleware calls the following one).
//
// A chain can be safely reused by calling Then() several times.
//     chain := server.NewChain(ratelimitHandler, csrfHandler)
//     indexPipe = chain.Then(indexHandler)
//     authPipe  = chain.Then(authHandler)
//
// Note that middlewares are called on every call to Then()
// and thus several instances of the same middleware will be created
// when a chain is reused in this previous example.
// For proper middleware, this should cause no problems.
//
// Then() treats nil as http.DefaultServeMux.
func (c Chain) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := range c {
		h = c[len(c)-1-i](h)
	}

	return h
}

// ThenFunc works identically to Then, but takes
// a HandlerFunc instead of a Handler.
//
// The following two statements are equivalent:
//     c.Then(http.HandlerFunc(fn))
//     c.ThenFunc(fn)
//
// ThenFunc provides all the guarantees of Then.
func (c Chain) ThenFunc(fn http.HandlerFunc) http.Handler {
	if fn == nil {
		return c.Then(nil)
	}

	return c.Then(fn)
}
