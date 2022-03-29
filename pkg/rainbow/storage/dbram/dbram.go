// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package dbram

import (
	"strings"

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

		db.insertOption(o)
	}

	return nil
}

func (db *DB) insertOption(o rainbow.Option) {
	for i, oldOpt := range db.optionsByProvider[o.Provider] {
		if oldOpt.Name == o.Name &&
			oldOpt.Type == o.Type &&
			oldOpt.Expiry == o.Expiry &&
			oldOpt.Strike == o.Strike {
			// Update the option attributes
			db.optionsByProvider[o.Provider][i] = o

			return
		}
	}

	// Append the new option
	db.optionsByProvider[o.Provider] = append(db.optionsByProvider[o.Provider], o)
}

func (db *DB) Options(args rainbow.StoreArgs) ([]rainbow.Option, error) {
	n := 0
	for _, o := range db.optionsByProvider {
		n += len(o)
	}

	options := make([]rainbow.Option, 0, n)
	for _, o := range db.optionsByProvider {
		options = append(options, o...)
	}

	filtered := make([]rainbow.Option, 0, n)

	for _, o := range options {
		if len(args.Asset) > 0 {
			if !containsAsset(o.Asset, args.Asset) {
				continue
			}
		}

		if len(args.Providers) > 0 {
			if !contains(o.Provider, args.Providers) {
				continue
			}
		}

		filtered = append(filtered, o)
	}

	return filtered, nil
}

// TODO go1.18: use generics with constraint comparable.
func containsAsset(elem string, arr []string) bool {
	for _, x := range arr {
		// use of strings.Contains because we allow some cases like WETH as ETH
		// TODO: sanitize data before put it in db
		if strings.Contains(elem, x) {
			return true
		}
	}

	return false
}

func contains(elem string, arr []string) bool {
	for _, x := range arr {
		if x == elem {
			return true
		}
	}

	return false
}
