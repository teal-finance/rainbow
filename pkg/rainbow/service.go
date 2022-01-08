// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"log"
	"time"
)

type Provider interface {
	Name() string
	Options() ([]Option, error)
}

type Service struct {
	providers []Provider
	store     Store
}

func NewService(p []Provider, s Store) Service {
	service := Service{
		providers: p,
		store:     s,
	}

	return service
}

// Run periodically fetch data from providers API and stores it in DB.
func (s *Service) Run() {
	ticker := time.NewTicker(10 * time.Minute)

	for ; true; <-ticker.C {
		s.FetchOptionsFromProviders()
	}
}

func (s *Service) FetchOptionsFromProviders() {
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
		log.Printf("Fetched %v=%v", p.Name(), len(o))
	}
}

func (s *Service) Options(args StoreArgs) ([]Option, error) {
	return s.store.Options(args)
}
