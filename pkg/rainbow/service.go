// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"log"
	"net/http"
	"time"
)

type Provider interface {
	Options() ([]Option, error)
}

type Store interface {
	InsertOptions(options []Option) error
	GetAllOptions() ([]Option, error)
}

type Service struct {
	provider Provider
	store    Store
	cache    Cache
}

func NewService(p Provider, s Store) Service {
	service := Service{
		provider: p,
		store:    s,
		cache:    Cache{},
	}

	service.initCache()

	return service
}

func (s *Service) initCache() {
	options, err := s.store.GetAllOptions()
	if err != nil {
		log.Print("ERROR store.GetAllOptions ", err)
	} else {
		s.cache.Refresh(options)
	}
}

func (s *Service) Handler() http.Handler {
	if s.cache.Empty() {
		s.initCache()
	}

	h := handler{c: &s.cache}

	return h.router()
}

// Run periodically gets and stores data from providers.
func (s *Service) Run() {
	ticker := time.NewTicker(10 * time.Minute)
	for ; true; <-ticker.C {
		options, err := s.OptionsFromProviders()
		if err != nil {
			log.Print("ERROR options from providers : ", err)
			continue // do not erase previously valid data (options, expiries, tables)
		}

		s.cache.Refresh(options)

		_ = s.store.InsertOptions(options)

		log.Print("Fetch options=", len(options))
	}
}

func (s *Service) OptionsFromProviders() ([]Option, error) {
	return s.provider.Options()
}

func (s *Service) Options() ([]Option, error) {
	return s.store.GetAllOptions()
}
