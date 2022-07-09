// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package api

import (
	"reflect"
	"testing"
)

func Test_specialValues(t *testing.T) {
	cases := []struct {
		name       string
		format     string
		values     []string
		wantFormat string
		wantValues []string
	}{
		{
			name:       "Keep all",
			format:     "",
			values:     []string{"ETH", "BTC", "ETH", "BTC"},
			wantFormat: "",
			wantValues: []string{"ETH", "BTC", "ETH", "BTC"},
		},
		{
			name:       "Keep two currencies",
			format:     "",
			values:     []string{"", "ALL", "csv", "BTC", "", "ETH"},
			wantFormat: "csv",
			wantValues: []string{"BTC", "ETH"},
		},
		{
			name:       "Keep the last format",
			format:     "",
			values:     []string{"", "ALL", "csv", "json", "", "tsv", "csv", "json", "tsv", "ALL"},
			wantFormat: "tsv",
			wantValues: nil,
		},
		{
			name:       "Input is nil",
			format:     "csv",
			values:     nil,
			wantFormat: "csv",
			wantValues: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			specialValues(&c.format, &c.values)

			if c.format != c.wantFormat {
				t.Errorf("format: got=%v want=%v", c.format, c.wantFormat)
			}

			if !reflect.DeepEqual(c.values, c.wantValues) {
				t.Errorf("values: got=%v want=%v", c.values, c.wantValues)
			}
		})
	}
}
