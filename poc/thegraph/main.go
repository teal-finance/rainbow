// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/teal-finance/rainbow/pkg/provider/zerox"
)

func main() {
	options := zerox.QueryTheGraph()
	spew.Dump(options)

}
