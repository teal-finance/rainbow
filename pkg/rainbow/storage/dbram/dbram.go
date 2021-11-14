// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package dbram

import (
	"fmt"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// DB client.
type DB struct {
	// optionsByProvider separates the options by provider
	optionsByProvider map[string][]rainbow.Option
}

func (db *DB) InsertOptions(p string, options []rainbow.Option) error {
	if db.optionsByProvider == nil {
		db.optionsByProvider = map[string][]rainbow.Option{p: options}
	} else {
		db.optionsByProvider[p] = options
	}

	return nil
}

func (db *DB) GetOptions(p string) ([]rainbow.Option, error) {
	options, ok := db.optionsByProvider[p]
	if !ok {
		return nil, fmt.Errorf("No data for %q", p)
	}

	return options, nil
}

func (db *DB) GetAllOptions() ([]rainbow.Option, error) {
	n := 0
	for _, o := range db.optionsByProvider {
		n += len(o)
	}

	options := make([]rainbow.Option, 0, n)
	for _, o := range db.optionsByProvider {
		options = append(options, o...)
	}

	return options, nil
}
