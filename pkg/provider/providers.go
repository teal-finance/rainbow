// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package provider

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

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
		deltaexchange.Provider{},
		deribit.Provider{},
		lyra.Provider{},
		psyoptions.Provider{},
		zerox.Provider{},
		zetamarkets.Provider{},
	}
}

// AllProvidersWithAlert returns all active providers with an alerter on anomalies.
// Do not panic if alerter endpoint is not reachable.
func AllProvidersWithAlert(o Oracle) []rainbow.Provider {
	hello := ":wave: Hi, RainbowOracle has just started"

	host, err := os.Hostname()
	if err == nil {
		hello += " on " + host
	}

	o.Query(hello)

	providers := AllProviders()
	for i, p := range providers {
		providers[i] = alerter{p, o}
	}

	return providers
}

type alerter struct {
	provider rainbow.Provider
	oracle   Oracle
}

func (a alerter) Name() string {
	return a.provider.Name()
}

func (a alerter) Options() ([]rainbow.Option, error) {
	options, err := a.provider.Options()

	go func() {
		if err != nil {
			a.oracle.Query(fmt.Sprintf(":alert: **%s**: API error: %s\n", a.Name(), err))
			return
		}

		if len(options) == 0 {
			a.oracle.Query(fmt.Sprintf(":question: **%s**: no options\n", a.Name()))
			return
		}
		// TODO: Check other anomalies
	}()

	return options, err
}

type Oracle struct {
	endpoint string
}

func NewOracle(endpoint string) Oracle {
	return Oracle{endpoint}
}

func (o Oracle) Query(query string) {
	resp, err := http.Post(
		o.endpoint,
		"application/json",
		bytes.NewBuffer([]byte(`{"username":"Oracle","text":"`+query+`"}`)),
	)
	if err != nil {
		log.Print("ERR Alerter POST: ", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		host := ""
		if url, er := url.Parse(o.endpoint); er == nil {
			host = url.Hostname()
		} else {
			log.Printf("ERR Alerter: URL=%q %s", o.endpoint, er)
		}
		log.Printf("ERR Alerter: %s returned status code %d", host, resp.StatusCode)
	}

	err = resp.Body.Close()
	if err != nil {
		log.Print("ERR Alerter: ", err)
	}
}
