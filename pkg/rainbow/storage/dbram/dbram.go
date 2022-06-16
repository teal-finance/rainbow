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
	for i := range options {
		provider := options[i].Provider
		if _, ok := db.optionsByProvider[provider]; !ok {
			db.optionsByProvider[provider] = []rainbow.Option{}
		}

		db.insertOption(&options[i])
	}
	return nil
}

func (db *DB) insertOption(newOpt *rainbow.Option) {
	options := db.optionsByProvider[newOpt.Provider]
	for i := range options {
		// if newOpt is already present, just update the fields and return
		if options[i].Name == newOpt.Name &&
			options[i].Type == newOpt.Type &&
			options[i].Expiry == newOpt.Expiry &&
			options[i].Strike == newOpt.Strike {
			options[i] = *newOpt
			return
		}
	}

	// Append the new option
	db.optionsByProvider[newOpt.Provider] = append(
		db.optionsByProvider[newOpt.Provider],
		*newOpt)
}

func (db *DB) Options(args rainbow.StoreArgs) ([]rainbow.Option, error) {
	total := db.numberOfOptions(args)
	filtered := make([]rainbow.Option, 0, total)

	for provider, options := range db.optionsByProvider {
		if !in(provider, args.Providers) {
			continue
		}

		for i := range options {
			if !contains(options[i].Asset, args.Assets) ||
				!startsWith(options[i].Expiry, args.Expiries) {
				continue
			}
			filtered = append(filtered, options[i])
		}
	}

	return filtered, nil
}

func (db *DB) numberOfOptions(args rainbow.StoreArgs) int {
	total := 0

	for provider, options := range db.optionsByProvider {
		if !in(provider, args.Providers) {
			continue
		}

		for i := range options {
			if !contains(options[i].Asset, args.Assets) ||
				!startsWith(options[i].Expiry, args.Expiries) {
				continue
			}
			total++
		}
	}

	return total
}

func in(provider string, wanted []string) bool {
	if len(wanted) == 0 {
		return true
	}

	for _, w := range wanted {
		if w == provider {
			return true
		}
	}
	return false
}

// TODO go1.18: use generics with constraint comparable.
func contains(asset string, subStrings []string) bool {
	if len(subStrings) == 0 {
		return true
	}

	for _, substr := range subStrings {
		// use of strings.Contains because we allow some cases like WETH as ETH
		// TODO: sanitize data before put it in db
		if strings.Contains(asset, substr) {
			return true
		}
	}
	return false
}

func startsWith(expiry string, prefixes []string) bool {
	if len(prefixes) == 0 {
		return true
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(expiry, prefix) {
			return true
		}
	}
	return false
}
