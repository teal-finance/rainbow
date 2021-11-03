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
