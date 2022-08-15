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

const (
	// useLoggerNotifier when no Mattermost endpoint.
	useLoggerNotifier = true
)

// AllProvidersNoAlerter returns all active providers without alerter.
func AllProvidersNoAlerter() []rainbow.Provider {
	return []rainbow.Provider{
		psyoptions.Provider{},    // separate psy and zeta to not
		deribit.Provider{},       //                  |
		lyra.Provider{},          //                  |
		zerox.Provider{},         //                  |
		zetamarkets.Provider{},   // <----------------` exhaust solana/serum rpc quota
		thales.Provider{},        // exotic market only live on /exotic
		deltaexchange.Provider{}, // last because slow (rate limit)
	}
}

// AllProvidersWithAlert returns all active providers with an alerter on anomalies.
// Do not panic if alerter endpoint is not reachable.
func AllProvidersWithAlert(n notifier.Notifier, namespace string) []rainbow.Provider {
	list := ""
	providers := AllProvidersNoAlerter()
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

// AllProviders returns all active providers with or without alerter
// depending on endpoint emptiness and on onlyMattermost.
func AllProviders(endpoint, namespace string) []rainbow.Provider {
	var n notifier.Notifier
	if endpoint != "" {
		n = mattermost.NewNotifier(endpoint)
	} else if useLoggerNotifier {
		n = logger.NewNotifier()
	}

	if n == nil {
		return AllProvidersNoAlerter()
	}
	return AllProvidersWithAlert(n, namespace)
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
