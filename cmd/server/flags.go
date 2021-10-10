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
)

const version = "Rainbow-1.0"

// exportPort is for Prometheus.
const exportPort = "no"

var (
	dev             = flag.Bool("dev", false, "Run rainbow in dev. mode")
	port            = flag.String("port", envStr("R_API_PORT", "no"), "API port, overseeds env. var. R_API_PORT")
	maxReqPerMinute = flag.Int("max-rate", envInt("R_MAX_REQ_PER_MINUTE", 6), "Max requests per minute, overseeds env. var. R_MAX_REQ_PER_MINUTE")
	maxReqBurst     = flag.Int("max-burst", envInt("R_MAX_REQ_BURST", 1), "Max requests during a burst, overseeds env. var. R_MAX_REQ_BURST")
)

func parseFlags() {
	flag.Parse()
	logFlags()
}

func logFlags() {
	log.Print("Dev. mode             -dev       = ", *dev)
	log.Print("R_API_PORT            -port      = ", *port)
	log.Print("R_MAX_REQ_PER_MINUTE  -max-rate  = ", *maxReqPerMinute)
	log.Print("R_MAX_REQ_BURST       -max-burst = ", *maxReqBurst)
}

func envStr(varName, defaultValue string) string {
	if value, ok := os.LookupEnv(varName); ok {
		return value
	}

	return defaultValue
}

func envInt(varName string, defaultValue int) int {
	if str, ok := os.LookupEnv(varName); ok {
		val, err := strconv.Atoi(str)
		if err != nil {
			return val
		}

		log.Printf("ERROR: Want integer but got %v=%v err: %v", varName, str, err)
	}

	return defaultValue
}
