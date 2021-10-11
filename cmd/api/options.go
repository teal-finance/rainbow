// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/teal-finance/rainbow"
	"github.com/teal-finance/rainbow/pkg/all"
)

var options *[]rainbow.Option

func alwaysCollectOptions() {
	for {
		newOptions, err := all.OptionsFromAllProviders()
		if err == nil {
			options = &newOptions
		}

		time.Sleep(10 * time.Minute)
	}
}

func replyOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(options); err != nil {
		log.Print("ERROR JSON Encode: ", err)
	}
}
