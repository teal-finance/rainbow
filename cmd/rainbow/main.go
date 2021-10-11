// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"github.com/teal-finance/rainbow/pkg/server"
)

func main() {
	parseFlags()

	if *apiPort <= 0 {
		printPrettyTabble()
		return
	}

	go alwaysCollectOptions()

	s := server.Server{
		Version:         version,
		DocURL:          "https://rainbow.teal.finance/doc",
		DevMode:         *dev,
		APIPort:         *apiPort,
		ExpPort:         *expPort,
		MaxReqBurst:     *maxReqBurst,
		MaxReqPerMinute: *maxReqPerMinute,
	}

	h := apiHandler(&s)

	s.RunServer(h)
}
