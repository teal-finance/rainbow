// Copyright (c) 2021 Teal.Finance
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

const (
	defaultAddr = "http://localhost"
	defaultPort = 8090
)

var (
	dev          = flag.Bool("dev", false, "Enable the developer mode")
	mainAddr     = flag.String("addr", envStr("MAIN_ADDR", defaultAddr), "Schema and DNS used for doc URL and CORS, has precedence over MAIN_ADDR")
	mainPort     = flag.Int("port", envInt("MAIN_PORT", defaultPort), "API port, has precedence over MAIN_PORT")
	expPort      = flag.Int("exp", envInt("EXP_PORT", 0), "Export port for Prometheus, has precedence over EXP_PORT")
	reqPerMinute = flag.Int("rate", envInt("REQ_PER_MINUTE", 88), "Max requests per minute, has precedence over REQ_PER_MINUTE")
	reqBurst     = flag.Int("burst", envInt("REQ_BURST", 22), "Max requests during a burst, has precedence over REQ_BURST")
	wwwDir       = flag.String("www", envStr("WWW_DIR", "frontend/dist"), "Folder of the web static files, has precedence over WWW_DIR")
	flagAlert    = flag.String("alert", "", "Mattermost webhook endpoint to activate alerter")
	listenAddr   string
)

func parseFlags() {
	flag.Parse()

	listenAddr = ":" + strconv.Itoa(*mainPort)

	if !*dev && *mainAddr == defaultAddr && *mainPort == defaultPort {
		*dev = true

		log.Print("Enable -dev mode because -addr and -port are not used")
	}

	log.Print("Dev. mode      -dev   = ", *dev)
	log.Print("MAIN_ADDR      -addr  = ", *mainAddr)
	log.Print("MAIN_PORT      -port  = ", *mainPort)
	log.Print("EXP_PORT       -exp   = ", *expPort)
	log.Print("REQ_PER_MINUTE -rate  = ", *reqPerMinute)
	log.Print("REQ_BURST      -burst = ", *reqBurst)
	log.Print("WWW_DIR        -www   = ", *wwwDir)
}

func envStr(varName, defaultValue string) string {
	value, ok := os.LookupEnv(varName)
	if !ok {
		value = defaultValue
	}
	return value
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
