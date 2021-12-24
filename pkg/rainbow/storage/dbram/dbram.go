// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package dbram

import (
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// DB client.
type DB struct {
	// optionsByProvider separates the options by provider
	optionsByProvider map[string][]rainbow.Option
}

func NewDB() *DB {
	return &DB{
		optionsByProvider: map[string][]rainbow.Option{},
	}
}

func (db *DB) InsertOptions(options []rainbow.Option) error {
	for _, o := range options {
		if _, ok := db.optionsByProvider[o.Provider]; !ok {
			db.optionsByProvider[o.Provider] = []rainbow.Option{}
		}
		db.optionsByProvider[o.Provider] = append(db.optionsByProvider[o.Provider], o)
	}

	return nil
}

func (db *DB) OptionsByProvider(p string) ([]rainbow.Option, error) {
	return db.optionsByProvider[p], nil
}

func (db *DB) Options() ([]rainbow.Option, error) {
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
