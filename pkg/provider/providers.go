// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package provider

import (
	"log"
	"os"

	"github.com/teal-finance/garcon"
	"github.com/teal-finance/notifier"
	"github.com/teal-finance/notifier/logger"
	"github.com/teal-finance/notifier/mattermost"
	"github.com/teal-finance/rainbow/pkg/provider/deltaexchange"
	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/provider/lyra"
	"github.com/teal-finance/rainbow/pkg/provider/psyoptions"
	"github.com/teal-finance/rainbow/pkg/provider/thales"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// AllProviders returns all supported providers.
func AllProviders() []rainbow.Provider {
	return []rainbow.Provider{
		psyoptions.Provider{},    // separate psy and zeta to not
		&deribit.Provider{},      //                  |
		lyra.Provider{},          //                  |
		zerox.Provider{},         //                  |
		zetamarkets.Provider{},   // <----------------` exhaust solana/serum rpc quota
		thales.Provider{},        // exotic market only live on /exotic
		deltaexchange.Provider{}, // last because slow (rate limit)
	}
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
	var n notifier.Notifier
	var namespace string

	if len(alerterOptions) > 0 {
		namespace = alerterOptions[0]
		if len(alerterOptions) > 1 && alerterOptions[1] != "" {
			endpoint := alerterOptions[1]
			n = mattermost.NewNotifier(endpoint)
		} else {
			n = logger.NewNotifier()
		}
	}

	providers := selectProviders(names)

	if n == nil {
		return providers
	}
	return AddAlert(providers, n, namespace)
}

// AddAlert returns the providers with an alerter on anomalies.
// Do not panic if alerter endpoint is not reachable.
func AddAlert(providers []rainbow.Provider, n notifier.Notifier, namespace string) []rainbow.Provider {
	list := ""
	for i, p := range providers {
		list += "\n" + "1. " + p.Name()
		providers[i] = newAlerter(namespace, p, n)
	}

	err := notifyStartup(n, namespace, list)
	if err != nil {
		log.Print("ERR Alerter: ", err)
	}

	return providers
}

// notifyStartup is called only once to notify when Rainbow is started.
func notifyStartup(n notifier.Notifier, namespace, list string) error {
	msg := ":wave: Rainbow **" + namespace + "** " + garcon.Version("") + " has just started"

	host, err := os.Hostname()
	if err == nil {
		msg += " on " + host
	}

	msg += " with:" + list

	return n.Notify(msg)
}
