// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/teal-finance/rainbow/pkg/provider"
)

const (
	queryDuration = 40 * time.Second
	sleepDuration = 10 * time.Minute
)

type Service struct {
	optionsJSON  []byte
	lastModified string
	expires      string
}

func (s *Service) CollectOptionsIndefinitely() {
	s.optionsJSON = []byte(`{"error":"initializing"}`)
	loopDuration := queryDuration + sleepDuration
	now := time.Now().UTC()

	for {
		beginTime := now

		options, err := provider.OptionsFromAllProviders()
		if err == nil {
			if b, err := json.Marshal(options); err != nil {
				log.Print("ERROR JSON Encode: ", err)
			} else if !bytes.Equal(s.optionsJSON, b) {
				s.optionsJSON = b
				s.lastModified = beginTime.Format(http.TimeFormat)
			}
		}

		// try to estimate the next refresh time
		s.expires = time.Now().UTC().Add(loopDuration).Format(http.TimeFormat)

		time.Sleep(sleepDuration)

		now = time.Now().UTC()
		loopDuration += now.Sub(beginTime)
		loopDuration /= 2 // compute a basic average
	}
}

func (s *Service) ReplyOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if s.lastModified != "" {
		w.Header().Set("Last-Modified", s.lastModified)
		w.Header().Set("Expires", s.expires)
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(s.optionsJSON)))

	_, _ = w.Write(s.optionsJSON)
}
