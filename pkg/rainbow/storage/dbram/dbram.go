package dbram

import (
	"fmt"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// DB client.
type DB struct {
	options  []rainbow.Option
	expiries rainbow.Expiries
	tables   map[string]rainbow.Table
}

func (db *DB) InsertOptions(options []rainbow.Option) error {
	db.options = options
	return nil
}

func (db *DB) InsertExpiries(expiries rainbow.Expiries) error {
	db.expiries = expiries
	return nil
}

func (db *DB) InsertTables(tables map[string]rainbow.Table) error {
	db.tables = tables
	return nil
}

func (db *DB) GetAllOptions() ([]rainbow.Option, error) {
	return db.options, nil
}

func (db *DB) GetExpiries() (rainbow.Expiries, error) {
	return db.expiries, nil
}

func (db *DB) GetTable(key string) (rainbow.Table, error) {
	t, ok := db.tables[key]
	if ok {
		return t, nil
	}

	return t, fmt.Errorf("No data found")
}
