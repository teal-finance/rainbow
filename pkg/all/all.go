// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package all

import (
	"log"

	"github.com/teal-finance/rainbow"
	"github.com/teal-finance/rainbow/pkg/deribit"
	"github.com/teal-finance/rainbow/pkg/psyoptions"
	"github.com/teal-finance/rainbow/pkg/zerox"
)

func OptionsFromAllProviders() ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	// psy
	psy, err := psyoptions.Options()
	if err != nil {
		log.Print("psy error ", err)
		return nil, err
	}

	options = append(options, psy...)

	// opyn
	op, err := zerox.Options()
	if err != nil {
		log.Print("opyn error ", err)
		return nil, err
	}

	options = append(options, op...)

	der, err := deribit.Options()
	if err != nil {
		log.Print("der error ", err)
		return nil, err
	}

	return append(options, der...), nil
}
