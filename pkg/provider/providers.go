// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package provider

import (
	"os"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/garcon/gg"
	"github.com/teal-finance/rainbow/pkg/provider/aevo"
	"github.com/teal-finance/rainbow/pkg/provider/deltaexchange"
	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/provider/lyra"
	"github.com/teal-finance/rainbow/pkg/provider/lyrav2"
	"github.com/teal-finance/rainbow/pkg/provider/rysk"
	"github.com/teal-finance/rainbow/pkg/provider/sdx"
	"github.com/teal-finance/rainbow/pkg/provider/synquote"
	"github.com/teal-finance/rainbow/pkg/provider/thales"
	"github.com/teal-finance/rainbow/pkg/provider/thalex"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone("provider")

// AllProviders returns all supported providers.
func AllProviders() []rainbow.Provider {
	return []rainbow.Provider{
		&deribit.Provider{},
		&rysk.Provider{},
		lyra.Provider{},
		lyrav2.Provider{},
		synquote.Provider{},
		&thalex.Provider{},
		aevo.Provider{},
		&sdx.Provider{},
		// zetamarkets.Provider{},	// paused platform due to current market conditions
		thales.Provider{},        // Thales = exotic options -> https://teal.finance/rainbow/exotic
		deltaexchange.Provider{}, // last because slow (rate limit)
	}
	// PsyOptions and Opyn are currently deprecated:
	// psyoptions.Provider{}
	// zerox.Provider{}
}

// Names lists the names of all supported providers.
func Names() []string {
	providers := AllProviders()
	names := make([]string, 0, len(providers))
	for _, p := range providers {
		names = append(names, p.Name())
	}
	return names
}

func selectProviders(names []string) []rainbow.Provider {
	selected := make([]rainbow.Provider, 0, len(names))
	for _, n := range names {
		for _, p := range AllProviders() {
			if n == p.Name() {
				selected = append(selected, p)
			}
		}
	}
	return selected
}

// Select returns all active providers with or without alerter
// depending on endpoint emptiness and on onlyMattermost.
func Select(names []string, alerterOptions ...string) []rainbow.Provider {
	providers := selectProviders(names)
	namespace, notifier := parseArgs(alerterOptions...)
	if notifier == nil {
		return providers
	}
	return AddAlert(providers, notifier, namespace)
}

func parseArgs(alerterOptions ...string) (namespace string, _ gg.Notifier) {
	if len(alerterOptions) == 0 {
		return "", nil
	}

	namespace = alerterOptions[0]

	endpoint := ""
	if len(alerterOptions) > 1 {
		endpoint = alerterOptions[1]
	}

	return namespace, gg.NewNotifier(endpoint)
}

// AddAlert returns the providers with an alerter on anomalies.
// Do not panic if alerter endpoint is not reachable.
func AddAlert(providers []rainbow.Provider, n gg.Notifier, namespace string) []rainbow.Provider {
	list := ""
	for i, p := range providers {
		list += "\n" + "1. " + p.Name()
		providers[i] = newAlerter(namespace, p, n)
	}

	err := notifyStartup(n, namespace, list)
	if err != nil {
		log.Error("Alerter:", err)
	}

	return providers
}

// notifyStartup is called only once to notify when Rainbow is started.
func notifyStartup(n gg.Notifier, namespace, list string) error {
	msg := ":wave: Rainbow **" + namespace + "** " + garcon.Version("") + " has just started"

	host, err := os.Hostname()
	if err == nil {
		msg += " on " + host
	}

	msg += " with:" + list

	return n.Notify(msg)
}
