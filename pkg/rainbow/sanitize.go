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

// sanitizeAsset removes "W" (or "w") in "WETC" and "WBTC".
func sanitizeAsset(asset string) string {
	if len(asset) >= 4 && (asset[0] == 'W' || asset[0] == 'w') {
		return asset[1:]
	}

	return asset
}

// prettyDate converts "2021-12-31 23:59:59" into "Dec 31".
func sanitizeDate(date string) string {
	t, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		log.Printf("WARN prettyDate() cannot parse %q", date)
		return date
	}
<<<<<<< HEAD:pkg/rainbow/clean.go

	const (
		monthLayout = "Jan"
		dayLayout   = "_2"
	)

	b := make([]byte, 0, len(monthLayout)+len(narrowSp)+len(dayLayout))
	b = t.AppendFormat(b, monthLayout)
	b = append(b, narrowSp...)
	b = t.AppendFormat(b, dayLayout)

	return string(b)
=======
	return t.Format("Jan 02")
>>>>>>> main:pkg/rainbow/sanitize.go
}
