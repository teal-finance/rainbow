package sqlite

import (
	"github.com/teal-finance/rainbow/pkg/rainbow"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Option struct {
	gorm.Model
	Name          string
	Type          string
	Asset         string
	Expiry        string
	ExchangeType  string
	Chain         string
	Layer         string
	Provider      string
	QuoteCurrency string
	Strike        float64
	// Bid           []Order `json:"bid"`
	// Ask           []Order `json:"ask"`
}

func newOption(o rainbow.Option) *Option {
	return &Option{
		Name:          o.Name,
		Type:          o.Type,
		Asset:         o.Asset,
		Expiry:        o.Expiry,
		ExchangeType:  o.ExchangeType,
		Chain:         o.Chain,
		Layer:         o.Layer,
		Provider:      o.Provider,
		QuoteCurrency: o.QuoteCurrency,
		Strike:        o.Strike,
	}
}

func (o *Option) toOptions() rainbow.Option {
	return rainbow.Option{
		Name:          o.Name,
		Type:          o.Type,
		Asset:         o.Asset,
		Expiry:        o.Expiry,
		ExchangeType:  o.ExchangeType,
		Chain:         o.Chain,
		Layer:         o.Layer,
		Provider:      o.Provider,
		QuoteCurrency: o.QuoteCurrency,
		Strike:        o.Strike,
	}
}

type DB struct {
	*gorm.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Option{})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) InsertOptions(provider string, options []rainbow.Option) error {
	res := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Option{})
	if res.Error != nil {
		return res.Error
	}

	for _, o := range options {
		o.Provider = provider

		res := db.Create(newOption(o))
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (db *DB) GetOptions(provider string) ([]rainbow.Option, error) {
	options := []Option{}
	res := db.Where("provider = ?", provider).Find(&options)
	if res.Error != nil {
		return []rainbow.Option{}, res.Error
	}

	results := make([]rainbow.Option, len(options))
	for i := range options {
		results[i] = options[i].toOptions()
	}

	return results, res.Error
}

func (db *DB) GetAllOptions() ([]rainbow.Option, error) {
	options := []Option{}
	res := db.Find(&options)
	if res.Error != nil {
		return []rainbow.Option{}, res.Error
	}

	results := make([]rainbow.Option, len(options))
	for i := range options {
		results[i] = options[i].toOptions()
	}

	return results, res.Error
}
