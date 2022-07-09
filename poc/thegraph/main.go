// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"github.com/spewerspew/spew"

	"github.com/teal-finance/rainbow/pkg/provider/zerox"
)

func main() {
	options := zerox.QueryTheGraph()
	// zerox.QueryTheGraph()
	spew.Dump(options)
}
