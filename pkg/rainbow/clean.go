// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

// cleanAsset removes "W" (or "w") in "WETC" and "WBTC".
func cleanAsset(asset string) string {
	if len(asset) >= 4 && (asset[0] == 'W' || asset[0] == 'w') {
		return asset[1:]
	}

	return asset
}
