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

type Service struct {
	provider Provider
	store    Store

	lastUpdate time.Time
}

func NewService(p Provider, s Store) Service {
	return Service{
		provider: p,
		store:    s,
	}
}

type Provider interface {
	Options() ([]Option, error)
}

type Store interface {
	InsertOptions(options []Option) error
	GetAllOptions() ([]Option, error)
}

// Run periodically gets and stores datas from providers.
func (s *Service) Run() {
	ticker := time.NewTicker(10 * time.Minute)
	for ; true; <-ticker.C {
		o, err := s.OptionsFromProviders()
		if err != nil {
			log.Print("ERROR options from providers : ", err)
		}
		log.Println("Store options from providers")
		s.lastUpdate = time.Now()
		s.store.InsertOptions(o)
	}
}

func (s *Service) OptionsFromProviders() ([]Option, error) {
	options, err := s.provider.Options()
	if err != nil {
		return []Option{}, err
	}
	return options, nil
}

func (s *Service) Options() ([]Option, error) {
	return s.store.GetAllOptions()
}
