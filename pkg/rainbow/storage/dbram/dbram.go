package dbram

import (
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// DB client.
type DB struct {
	options  []rainbow.Option
	cpFormat rainbow.CPFormat
}

func (db *DB) InsertOptions(options []rainbow.Option) error {
	db.options = options
	return nil
}

func (db *DB) InsertCPFormat(table rainbow.CPFormat) error {
	db.cpFormat = table
	return nil
}

func (db *DB) GetAllOptions() ([]rainbow.Option, error) {
	return db.options, nil
}

func (db *DB) GetCPFormat() (rainbow.CPFormat, error) {
	return db.cpFormat, nil
}
