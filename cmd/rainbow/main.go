// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/teal-finance/rainbow/pkg/server"
)

const version = "Rainbow-0.1"

var (
	dev             = flag.Bool("dev", false, "Run rainbow in dev. mode")
	apiPort         = flag.Int("port", envInt("API_PORT", 0), "API port, overseeds env. var. API_PORT")
	expPort         = flag.Int("exp", envInt("EXP_PORT", 0), "Export port for Prometheus, overseeds env. var. EXP_PORT")
	maxReqPerMinute = flag.Int("max-rate", envInt("MAX_REQ_PER_MINUTE", 6), "Max requests per minute, overseeds env. var. MAX_REQ_PER_MINUTE")
	maxReqBurst     = flag.Int("max-burst", envInt("MAX_REQ_BURST", 1), "Max requests during a burst, overseeds env. var. MAX_REQ_BURST")
)

func main() {
	flag.Parse()
	logFlags()

	if *apiPort <= 0 {
		printPrettyTable()
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
		OPAFiles:        nil,
	}

	h := apiHandler(&s)

	if err := s.RunServer(h); err != nil {
		log.Fatal(err)
	}
}

func logFlags() {
	log.Print("Dev. mode           -dev       = ", *dev)
	log.Print("API_PORT            -port      = ", *apiPort)
	log.Print("EXP_PORT            -exp       = ", *expPort)
	log.Print("MAX_REQ_PER_MINUTE  -max-rate  = ", *maxReqPerMinute)
	log.Print("MAX_REQ_BURST       -max-burst = ", *maxReqBurst)
}

func envStr(varName, defaultValue string) string {
	if value, ok := os.LookupEnv(varName); ok {
		return value
	}

	return defaultValue
}

func envInt(varName string, defaultValue int) int {
	if str, ok := os.LookupEnv(varName); ok {
		value, err := strconv.Atoi(str)
		if err == nil {
			return value
		}

		log.Fatalf("ERROR: Want integer but got %v=%v err: %v", varName, str, err)
	}

	return defaultValue
}
