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
	Name() string
	Options() ([]Option, error)
}

type Service struct {
	providers []Provider
	store     Store
	cache     Cache
}

func NewService(p []Provider, s Store) Service {
	service := Service{
		providers: p,
		store:     s,
		cache:     Cache{},
	}

	service.initCache()

	return service
}

func (s *Service) initCache() {
	if s.cache.Empty() {
		options, err := s.store.Options(StoreArgs{})
		if err != nil {
			log.Print("ERROR store.GetAllOptions ", err)
			return
		}

		s.cache.Refresh(options)
	}
}

func (s *Service) Handler() http.Handler {
	s.initCache()

	h := handler{c: &s.cache}

	return h.router()
}

// Run periodically fetch data from providers API and stores it in DB.
func (s *Service) Run() {
	ticker := time.NewTicker(10 * time.Minute)

	for ; true; <-ticker.C {
		options := s.OptionsFromProviders()
		s.cache.Refresh(options)
	}
}

func (s *Service) OptionsFromProviders() []Option {
	var options []Option

	for _, p := range s.providers {
		o, err := p.Options()
		if err != nil {
			log.Print("WARN fetching data from ", p, " : ", err)
			continue
		}

		err = s.store.InsertOptions(o)
		if err != nil {
			log.Print("WARN cannot store data in DB for ", p, " : ", err)
			continue
		}

		options = append(options, o...)
		log.Printf("Fetched %v=%v", p.Name(), len(o))
	}

	return options
}

func (s *Service) Options(args StoreArgs) ([]Option, error) {
	return s.store.Options(args)
}

func (s *Service) CallPut() CallPut {
	return s.cache.callPut
}
