// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package provider

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/teal-finance/rainbow/pkg/provider/deltaexchange"
	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/provider/lyra"
	"github.com/teal-finance/rainbow/pkg/provider/psyoptions"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// AllProviders returns all active providers.
func AllProviders() []rainbow.Provider {
	// changing the order to not exhaust our solana/serum rpc quota
	// used by zeta and psy
	return []rainbow.Provider{
		//zetamarkets.Provider{},
		//zerox.Provider{},
		deribit.Provider{},
		lyra.Provider{},
		deltaexchange.Provider{},
		//psyoptions.Provider{},
	}
}

// AllProvidersWithAlert returns all active providers with an alerter on anormalities.
func AllProvidersWithAlert(o Oracle) []rainbow.Provider {
	err := o.Query("Hello, I am the new updated Oracle !")
	if err != nil {
		panic(err)
	}

	// changing the order to not exhaust our solana/serum rpc quota
	// used by zeta and psy
	return []rainbow.Provider{
		withAlert{zetamarkets.Provider{}, o},
		withAlert{zerox.Provider{}, o},
		withAlert{deribit.Provider{}, o},
		withAlert{lyra.Provider{}, o},
		withAlert{deltaexchange.Provider{}, o},
		withAlert{psyoptions.Provider{}, o},
	}
}

type withAlert struct {
	prov   rainbow.Provider
	oracle Oracle
}

func (p withAlert) Name() string {
	return p.prov.Name()
}

func (p withAlert) Options() ([]rainbow.Option, error) {
	options, err := p.prov.Options()

	go func() {
		if err != nil {
			p.oracle.Query(fmt.Sprintf(":alert: **%s**: api error: %s\n", p.Name(), err))
			return
		}

		// TO DO other anomality checks
	}()

	return options, err
}

type Oracle struct {
	endpoint string
}

func NewOracle(endpoint string) Oracle {
	return Oracle{endpoint}
}

func (o Oracle) Query(query string) error {
	resp, err := http.Post(
		o.endpoint,
		"application/json",
		bytes.NewBuffer([]byte(`{"username":"Oracle","text":"`+query+`"}`)),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("framateam.org returned status code %d", resp.StatusCode)
	}
	return nil
}
