// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"log"

	"github.com/teal-finance/rainbow/pkg/opyn"
)

func main() {
	options := opyn.Instruments()
	log.Print("Options: ", options)
}
