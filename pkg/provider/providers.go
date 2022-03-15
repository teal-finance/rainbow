// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package provider

import (
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
	//changing the order to not exhaust our solana/serum rpc quota
	//used by zeta and psy
	return []rainbow.Provider{
		zetamarkets.Provider{},
		zerox.Provider{},
		deribit.Provider{},
		lyra.Provider{},
		deltaexchange.Provider{},
		psyoptions.Provider{},
	}
}
