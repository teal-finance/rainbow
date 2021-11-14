// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type Provider interface {
	Name() string
	Options() ([]Option, error)
}

type Store interface {
	InsertOptions(provider string, options []Option) error
	GetOptions(provider string) ([]Option, error)
	GetAllOptions() ([]Option, error)
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
		options, err := s.store.GetAllOptions()
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

	stat := ""

	for _, p := range s.providers {
		o, err := p.Options()
		if err != nil {
			log.Print("WARN fetching data from ", p, " : ", err)

			o, err = s.store.GetOptions(p.Name())
			if err != nil {
				log.Print("WARN no data in DB for ", p, " : ", err)
				continue
			}
		}

		err = s.store.InsertOptions(p.Name(), o)
		if err != nil {
			log.Print("WARN cannot store data in DB for ", p, " : ", err)
		}

		stat += " " + p.Name() + "=" + strconv.Itoa(len(o))
		options = append(options, o...)
	}

	if stat != "" {
		log.Print("Fetched" + stat)
	}

	return options
}

func (s *Service) Options() ([]Option, error) {
	return s.store.GetAllOptions()
}
