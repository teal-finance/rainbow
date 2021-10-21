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
	"strings"
)

var (
	dev             = flag.Bool("dev", false, "Run rainbow in dev. mode")
	mainAddr        = flag.String("addr", envStr("MAIN_ADDR", "http://localhost"), "Schema and DNS used for doc URL and CORS, has precedence over MAIN_ADDR")
	mainPort        = flag.Int("port", envInt("MAIN_PORT", 1234), "API port, has precedence over MAIN_PORT")
	expPort         = flag.Int("exp", envInt("EXP_PORT", 0), "Export port for Prometheus, has precedence over EXP_PORT")
	maxReqPerMinute = flag.Int("rate", envInt("REQ_PER_MINUTE", 30), "Max requests per minute, has precedence over REQ_PER_MINUTE")
	maxReqBurst     = flag.Int("burst", envInt("REQ_BURST", 10), "Max requests during a burst, has precedence over REQ_BURST")
	wwwDir          = flag.String("www", envStr("WWW_DIR", "frontend/dist"), "Folder of the web static files, has precedence over WWW_DIR")
	opaFlag         = flag.String("opa", "", "Policy files (comma-separated filenames) for the Open Policy Agent using the Datalog/Rego format")
	opaFilenames    []string

	listenAddr string
)

func parseFlags() {
	flag.Parse()

	if *opaFlag != "" {
		opaFilenames = strings.Split(*opaFlag, ",")
	}

	listenAddr = ":" + strconv.Itoa(*mainPort)

	log.Print("Dev. mode      -dev   = ", *dev)
	log.Print("MAIN_ADDR      -addr  = ", *mainAddr)
	log.Print("MAIN_PORT      -port  = ", *mainPort)
	log.Print("EXP_PORT       -exp   = ", *expPort)
	log.Print("REQ_PER_MINUTE -rate  = ", *maxReqPerMinute)
	log.Print("REQ_BURST      -burst = ", *maxReqBurst)
	log.Print("WWW_DIR        -www   = ", *wwwDir)
	log.Printf("Policy files   -opa   = #%v %q", len(opaFilenames), opaFilenames)
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
