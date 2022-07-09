// Copyright 2021-2022 Teal.Finance contributors
// This file is part of Teal.Finance/Rainbow,
// a screener fo DeFi options, under the MIT License.
// SPDX-License-Identifier: MIT

package provider

import (
	"fmt"
	"log"
	"os"

	"github.com/teal-finance/notifier"
	"github.com/teal-finance/rainbow/pkg/provider/deltaexchange"
	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/provider/lyra"
	"github.com/teal-finance/rainbow/pkg/provider/psyoptions"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// AllProviders returns all active providers (without alerter).
func AllProviders() []rainbow.Provider {
	// changing the order to not exhaust our solana/serum rpc quota
	// used by zeta and psy
	return []rainbow.Provider{
		deribit.Provider{},
		lyra.Provider{},
		psyoptions.Provider{},
		zerox.Provider{},
		zetamarkets.Provider{},
		deltaexchange.Provider{}, // last because slow (rate limit)
	}
}

// AllProvidersWithAlert returns all active providers with an alerter on anomalies.
// Do not panic if alerter endpoint is not reachable.
func AllProvidersWithAlert(n notifier.Notifier) []rainbow.Provider {
	hello := ":wave: Hi, RainbowAlerter has just started"

	host, err := os.Hostname()
	if err == nil {
		hello += " on " + host
	}

	_ = n.Notify(hello)

	providers := AllProviders()
	for i, p := range providers {
		providers[i] = alerter{p, n}
	}

	return providers
}

type alerter struct {
	provider rainbow.Provider
	notifier notifier.Notifier
}

func (a alerter) Name() string {
	return a.provider.Name()
}

func (a alerter) Options() ([]rainbow.Option, error) {
	options, err := a.provider.Options()

	go func() {
		err := a.vet(options, err)
		if err != nil {
			log.Print("ERR Alerter: ", err)
		}
	}()

	return options, err
}

func (a alerter) vet(options []rainbow.Option, err error) error {
	if err != nil {
		return a.notifier.Notify(fmt.Sprintf(":alert: **%s**: API error: %s\n", a.Name(), err))
	}

	if len(options) == 0 && a.Name() != "Lyra" {
		return a.notifier.Notify(fmt.Sprintf(":question: **%s**: no options\n", a.Name()))
	}

	// TODO: Check other anomalies

	return nil
}
