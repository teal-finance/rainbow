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
	InsertCPFormat(CPFormat) error

	GetAllOptions() ([]Option, error)
	GetCPFormat() (CPFormat, error)
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

		cp := buildCPFormat(o)

		_ = s.store.InsertOptions(o)
		_ = s.store.InsertCPFormat(cp)

		log.Printf("Update options=%v rows=%v", len(o), len(cp.Rows))
	}
}

func (s *Service) OptionsFromProviders() ([]Option, error) {
	return s.provider.Options()
}

func (s *Service) Options() ([]Option, error) {
	return s.store.GetAllOptions()
}
