package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// DB client.
type DB struct {
	*sql.DB
}

func NewDB(file string) (*DB, error) {
	conn, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	db := &DB{DB: conn}

	err = db.initTables()
	if err != nil {
		return nil, err
	}

	return db, nil
}

const (
	orderAsk = "ask"
	orderBid = "bid"
)

func (db *DB) initTables() error {
	const req = `
		CREATE TABLE IF NOT EXISTS options (
			name TEXT NOT NULL PRIMARY KEY,
			type TEXT,
			asset TEXT,
			expiry TEXT,
			exchange TEXT,
			chain TEXT,
			layer TEXT,
			provider TEXT,
			currency TEXT,
			strike TEXT
		);
		CREATE TABLE IF NOT EXISTS orders(
			option TEXT NOT NULL,
			side TEXT NOT NULL,
			size REAL,
			price REAL,
			FOREIGN KEY(option) REFERENCES options(name)
		);
	`
	_, err := db.Exec(req)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) InsertOptions(options []rainbow.Option) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT OR REPLACE INTO options(
			name, type, asset, expiry, exchange, 
			chain, layer, provider, currency, strike
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmtOrder, err := tx.Prepare(`INSERT INTO orders(option, side, size, price) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmtOrder.Close()

	for _, o := range options {
		_, err = stmt.Exec(
			o.Name,
			o.Type,
			o.Asset,
			o.Expiry,
			o.ExchangeType,
			o.Chain,
			o.Layer,
			o.Provider,
			o.QuoteCurrency,
			o.Strike,
		)
		if err != nil {
			return err
		}

		for _, a := range o.Ask {
			_, err = stmtOrder.Exec(o.Name, orderAsk, a.Size, a.Price)
			if err != nil {
				return err
			}
		}

		for _, a := range o.Bid {
			_, err = stmtOrder.Exec(o.Name, orderBid, a.Size, a.Price)
			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Options(args rainbow.StoreArgs) ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	rows, err := db.Query(queryWithArgs(`SELECT * FROM options`, args))
	if err != nil {
		return []rainbow.Option{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var o rainbow.Option
		err = rows.Scan(
			&o.Name,
			&o.Type,
			&o.Asset,
			&o.Expiry,
			&o.ExchangeType,
			&o.Chain,
			&o.Layer,
			&o.Provider,
			&o.QuoteCurrency,
			&o.Strike)
		if err != nil {
			return []rainbow.Option{}, err
		}
		options = append(options, o)
	}

	for i := range options {
		rows, err := db.Query(`SELECT side, size, price FROM orders WHERE option = ?`, options[i].Name)
		if err != nil {
			return []rainbow.Option{}, err
		}
		defer rows.Close()
		for rows.Next() {
			var side string
			var o rainbow.Order
			err = rows.Scan(&side, &o.Size, &o.Price)
			if err != nil {
				return []rainbow.Option{}, err
			}
			switch side {
			case orderAsk:
				options[i].Ask = append(options[i].Ask, o)
			case orderBid:
				options[i].Bid = append(options[i].Bid, o)
			}
		}
	}

	return options, nil
}

func queryWithArgs(query string, args rainbow.StoreArgs) string {
	filterStr := ""
	where := []string{}

	if len(args.Providers) > 0 {
		whereProvider := []string{}
		for _, provider := range args.Providers {
			whereProvider = append(whereProvider, fmt.Sprintf("provider LIKE '%%%v%%'", provider))
		}
		where = append(where, fmt.Sprintf("(%v)", strings.Join(whereProvider, " OR ")))
	}

	if len(args.Asset) > 0 {
		whereAsset := []string{}
		for _, asset := range args.Asset {
			whereAsset = append(whereAsset, fmt.Sprintf("asset LIKE '%%%v%%'", asset))
		}
		where = append(where, fmt.Sprintf("(%v)", strings.Join(whereAsset, " OR ")))
	}

	if len(where) > 0 {
		filterStr = fmt.Sprintf("WHERE %s", strings.Join(where, " AND "))
	}

	query = fmt.Sprintf("SELECT * FROM (%v) AS query %v", query, filterStr)

	return query
}
