package rainbow

import (
	"log"
	"time"
)

var (
	options    = []Option{}
	lastUpdate time.Time
)

type Service struct {
	provider Provider
}

func NewService(p Provider) Service {
	return Service{
		provider: p,
	}
}

type Provider interface {
	Options() ([]Option, error)
}

// Run periodically gets and stores datas from providers.
func (s *Service) Run() {
	ticker := time.NewTicker(10 * time.Minute)
	for ; true; <-ticker.C {
		o, err := s.Options()
		if err != nil {
			log.Print("ERROR options from providers : ", err)
		}
		lastUpdate = time.Now()
		options = o
	}
}

func (s *Service) Options() ([]Option, error) {
	options, err := s.provider.Options()
	if err != nil {
		return []Option{}, err
	}

	return options, nil
}
