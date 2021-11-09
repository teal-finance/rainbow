// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package provider

import (
	"fmt"

	"github.com/teal-finance/rainbow/pkg/provider/deribit"
	"github.com/teal-finance/rainbow/pkg/provider/psyoptions"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// AllProvider gets data from all providers.
type AllProvider struct{}

func (p AllProvider) Options() ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	// psy
	psy, err := psyoptions.Options()
	if err != nil {
		return nil, fmt.Errorf("getting data from psy : %w", err)
	}

	options = append(options, psy...)

	// opyn
	op, err := zerox.Options()
	if err != nil {
		return nil, fmt.Errorf("getting data from opyn : %w", err)
	}

	options = append(options, op...)

	der, err := deribit.Options()
	if err != nil {
		return nil, fmt.Errorf("getting data from deribit : %w", err)
	}

	options = append(options, der...)

	return options, nil
}
