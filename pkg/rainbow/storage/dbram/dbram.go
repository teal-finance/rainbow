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
	options []rainbow.Option
}

func (db *DB) InsertOptions(options []rainbow.Option) error {
	db.options = options
	return nil
}

func (db *DB) GetAllOptions() ([]rainbow.Option, error) {
	return db.options, nil
}
