// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package rainbow

import (
	"time"

	"github.com/teal-finance/emo"
)

var log = emo.NewZone("srv")

type Provider interface {
	Name() string
	Options() ([]Option, error)
}

type Service struct {
	providers []Provider
	store     Store
}

func NewService(p []Provider, s Store) Service {
	return Service{
		providers: p,
		store:     s,
	}
}

// Run periodically fetch data from providers API and stores it in DB.
func (s *Service) Run(loopTime time.Duration) {
	ticker := time.NewTicker(loopTime)

	for ; true; <-ticker.C {
		s.FetchOptionsFromProviders()
	}
}

func (s *Service) FetchOptionsFromProviders() {
	for _, p := range s.providers {
		o, err := p.Options()
		if err != nil {
			log.Error(p.Name(), err)
			continue
		}

		log.Printf(p.Name()+": fetched %v options", len(o))

		err = s.store.InsertOptions(o)
		if err != nil {
			log.Error(p.Name(), "cannot store data:", err)
			continue
		}
	}
}

func (s *Service) Options(args StoreArgs) ([]Option, error) {
	return s.store.Options(args)
}
