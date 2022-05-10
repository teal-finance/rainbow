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
	dev          = flag.Bool("dev", false, "Enable the developer mode (enabled by default if -addr and -port are not used)")
	mainAddr     = flag.String("addr", envStr("MAIN_ADDR", defaultAddr), "Schema and DNS used for doc URL and CORS, has precedence over MAIN_ADDR")
	mainPort     = flag.Int("port", envInt("MAIN_PORT", defaultPort), "API port, has precedence over MAIN_PORT")
	expPort      = flag.Int("exp", envInt("EXP_PORT", 0), "Export port for Prometheus, has precedence over EXP_PORT")
	reqPerMinute = flag.Int("rate", envInt("REQ_PER_MINUTE", 88), "Max requests per minute, has precedence over REQ_PER_MINUTE")
	reqBurst     = flag.Int("burst", envInt("REQ_BURST", 22), "Max requests during a burst, has precedence over REQ_BURST")
	wwwDir       = flag.String("www", envStr("WWW_DIR", "frontend/dist"), "Folder of the web static files, has precedence over WWW_DIR")
	alert        = flag.String("alert", envStr("ALERT_URL", ""), "Webhook endpoint to notify anomalies, has precedence over ALERT_URL")
	aes          = flag.String("aes", envStr("AES_128", ""), " 128-bit AES key (32 hex chars) for the session cookies, has precedence over AES_128")
	hmac         = flag.String("hmac", envStr("HMAC_SHA256", ""), "HMAC-SHA256 key (64 hex chars) for the JWT tokens, has precedence over HMAC_SHA256")
	listenAddr   string
)

func parseFlags() {
	flag.Parse()

	listenAddr = ":" + strconv.Itoa(*mainPort)

	if !*dev && *mainAddr == defaultAddr && *mainPort == defaultPort {
		*dev = true
		log.Print("Enable -dev mode because -addr and -port flags are not used")
	}

	log.Print("Dev. mode      -dev   = ", *dev)
	log.Print("MAIN_ADDR      -addr  = ", *mainAddr)
	log.Print("MAIN_PORT      -port  = ", *mainPort)
	log.Print("EXP_PORT       -exp   = ", *expPort)
	log.Print("REQ_PER_MINUTE -rate  = ", *reqPerMinute)
	log.Print("REQ_BURST      -burst = ", *reqBurst)
	log.Print("WWW_DIR        -www   = ", *wwwDir)
	log.Print("ALERT_URL      -alert len=", len(*alert))
	log.Print("AES_128        -aes   len=", len(*aes), " (need 32 hexadecimal numbers)")
	log.Print("HMAC_SHA256    -hmac  len=", len(*hmac), " (need 16 hexadecimal numbers)")

	// mandatory: -aes or -hmac
	if len(*aes) == 0 && len(*hmac) == 0 {
		if *dev {
			*hmac = "9d2e0a02121179a3c3de1b035ae1355b1548781c8ce8538a1dc0853a12dfb13d"
		} else {
			log.Fatal("Missing secret key for the tokens (cookies). " +
				"Please provide it using the -aes or -hmac flag " +
				"(or the AES_128 or HMAC_SHA256 env. var.)")
		}
	}
	if len(*aes) > 0 && len(*hmac) > 0 {
		log.Print("WAR: Should use -aes or -hmac, not both in the same time")
	}
}

// envStr looks up the given key from the environment,
// returning its value if it exists, and otherwise
// returning the given fallback value.
func envStr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// envInt looks up the given key from the environment and expects an integer,
// returning the integer value if it exists, and otherwise returning the fallback value.
// If the environment variable has a value but it can't be parsed as an integer,
// envInt terminates the program.
func envInt(key string, fallback int) int {
	if s, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("ERR: want integer but got %v=%q err: %v", key, s, err)
		}
		return v
	}
	return fallback
}
