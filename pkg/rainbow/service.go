package rainbow

import (
	"log"
	"time"
)

type Service struct {
	provider Provider
	store    Store
}

func NewService(p Provider, s Store) Service {
	return Service{
		provider: p,
		store:    s,
	}
}

type Provider interface {
	Options() ([]Option, error)
}

type Store interface {
	InsertOptions(options []Option) error
	InsertExpiries(Expiries) error
	InsertTables(map[string]Table) error

	GetAllOptions() ([]Option, error)
	GetExpiries() (Expiries, error)
	GetTable(key string) (Table, error)
}

// Run periodically gets and stores data from providers.
func (s *Service) Run() {
	ticker := time.NewTicker(10 * time.Minute)
	for ; true; <-ticker.C {
		o, err := s.OptionsFromProviders()
		if err != nil {
			log.Print("ERROR options from providers : ", err)
			continue // do not erase previously valid data (options, expiries, tables)
		}

		e := buildExpiries(o)
		t := buildTables(o)

		_ = s.store.InsertOptions(o)
		_ = s.store.InsertExpiries(e)
		_ = s.store.InsertTables(t)

		log.Printf("Update options=%v expiries=%v tables=%v",
			len(o), len(e.AssetToExpiries), len(t))
	}
}

func (s *Service) OptionsFromProviders() ([]Option, error) {
	return s.provider.Options()
}

func (s *Service) Options() ([]Option, error) {
	return s.store.GetAllOptions()
}
