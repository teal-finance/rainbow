// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package provider

import (
	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/provider/psyoptions"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
	"github.com/teal-finance/rainbow/pkg/provider/zetamarkets"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// AllProviders returns all active providers.
func AllProviders() []rainbow.Provider {
	return []rainbow.Provider{
		psyoptions.Provider{},
		zetamarkets.Provider{},
		zerox.Provider{},
		deribit.Provider{},
	}
}
