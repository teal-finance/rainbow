// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"log"
	"time"
)

// cleanAsset removes "W" (or "w") in "WETC" and "WBTC".
func cleanAsset(asset string) string {
	if len(asset) >= 4 && (asset[0] == 'W' || asset[0] == 'w') {
		return asset[1:]
	}

	return asset
}

// prettyDate converts "2021-12-31 23:59:59" into "Dec 31".
func prettyDate(date string) string {
	t, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		log.Printf("WARN prettyDate() cannot parse %q", date)
		return date
	}

	return t.Format("Jan _2")
}
