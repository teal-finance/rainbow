# Teal.Finance/Server

![logo](logo.jpg) | <big>Opinionated boilerplate HTTP server with CORS, OPA, Prometheus, rate-limiter… for API and static website.</big>
-|-

## Origin

This library was originally developed as part of the project
[Rainbow](https://github.com/teal-finance/rainbow) during hackathons,
based on older Teal.Finance products,
and then moved to [its own repository](https://github.com/teal-finance/server).

## Features

Teal.Finance/Server supports:

* Metrics server exporting data to Prometheus or other monitoring services ;
* File server intended for static web files ;
* HTTP/REST server for API endpoints (compatible any Go-standard HTTP handlers) ;
* Chained middlewares (fork of github.com/justinas/alice)
* Auto-completed error response in JSON format ;
* Middleware: authentication rules based on Datalog/Rego files using [Open Policy Agent](https://www.openpolicyagent.org) ;
* Middleware: rate limiter to prevent flooding by incoming requests ;
* Middleware: logging of incoming requests ;
* Middleware: Cross-Origin Resource Sharing (CORS).

## License

[LGPL-3.0-or-later](https://spdx.org/licenses/LGPL-3.0-or-later.html):
GNU Lesser General Public License v3.0 or later
([tl;drLegal](https://tldrlegal.com/license/gnu-lesser-general-public-license-v3-(lgpl-3)),
[Choosealicense.com](https://choosealicense.com/licenses/lgpl-3.0/)).
See the [LICENSE](LICENSE) file.

Except the two example files under CC0-1.0 (Creative Commons Zero v1.0 Universal)
and the file [chain.go](chain/chain.go) (fork) under the MIT License.

## Examples

See also a complete real example in the repo
[github.com/teal-finance/rainbow](https://github.com/teal-finance/rainbow/blob/main/cmd/server/main.go).

## High-level

The following code uses the high-level function `Server.RunServer()`.

```go
package main

import (
    "log"

    "github.com/teal-finance/server"
)

func main() {
    s := server.Server{
        Version:        "MyApp-1.2.0",
        Resp:           "https://my.dns.co/doc",
        AllowedOrigins: []string{"https://my.dns.co"},
        OPAFilenames:   []string{"example-auth.rego"},
    }

    h := myHandler()

    // main port 8080, export port 9093, rate limiter 10 20, dev mode 
    log.Fatal(s.RunServer(h, 8080, 9093, 10, 20, true))
}
```

Run the [high-level example](examples/high-level/main.go):

```
$ go build ./examples/high-level && ./high-level
2021/10/19 23:41:50 Prometheus export http://localhost:9093
2021/10/19 23:41:50 CORS: Set origin: https://my.dns.co
2021/10/19 23:41:50 Middleware CORS: {AllowedOrigins:[] AllowOriginFunc:0x556b48e30960 AllowOriginRequestFunc:<nil> AllowedMethods:[GET] AllowedHeaders:[Origin Accept Content-Type Authorization Cookie] ExposedHeaders:[] MaxAge:60 AllowCredentials:true OptionsPassthrough:false Debug:true}
2021/10/19 23:41:50 OPA: load "example-auth.rego"
2021/10/19 23:41:50 Middleware OPA: map[example-auth.rego:package auth

default allow = false
tokens := {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dnZWRJbkFzIjoiYWRtaW4iLCJpYXQiOjE0MjI3Nzk2Mzh9.gzSraSYS8EXBxLN_oWnFSRgCzcmJmMjLiuyu5CSpyHI"} { true }
allow = true { __local0__ = input.token; data.auth.tokens[__local0__] }]
2021/10/19 23:41:50 Middleware response HTTP header: Set Server MyApp-1.2.0
2021/10/19 23:41:50 Middleware RateLimiter: burst=100 rate=5/s
2021/10/19 23:41:50 Middleware logger: log requested URLs and remote addresses
2021/10/19 23:41:50 Server listening on http://localhost:8080
```

Test the API endpoint with default `curl` HTTP headers:

```
$ curl -D - http://localhost:8080/api/v1/items
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Server: MyApp-1.2.0
Vary: Origin
X-Content-Type-Options: nosniff
Date: Tue, 19 Oct 2021 21:42:29 GMT
Content-Length: 77

{"error":"Unauthorized","doc":"https://my.dns.co/doc","path":"/api/v1/items"}
```

The corresponding server logs in debug mode:

```
2021/10/19 23:42:29 in  GET /api/v1/items [::1]:54796
[cors] 2021/10/19 23:42:29 Handler: Actual request
[cors] 2021/10/19 23:42:29   Actual request no headers added: missing origin
2021/10/19 23:42:29 OPA unauthorize [::1]:54796 /api/v1/items
2021/10/19 23:42:29 out GET /api/v1/items 337.751µs
```

Test the API endpoint with valid Authorization header:

```
$ curl -D - -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dnZWRJbkFzIjoiYWRtaW4iLCJpYXQiOjE0MjI3Nzk2Mzh9.gzSraSYS8EXBxLN_oWnFSRgCzcmJmMjLiuyu5CSpyHI' http://localhost:8080/api/v1/items
HTTP/1.1 200 OK
Content-Type: application/json
Server: MyApp-1.2.0
Vary: Origin
Date: Tue, 19 Oct 2021 21:43:55 GMT
Content-Length: 25

["item1","item2","item3"]
```

The corresponding server logs:

```
2021/10/19 23:43:55 in  GET /api/v1/items [::1]:54798
[cors] 2021/10/19 23:43:55 Handler: Actual request
[cors] 2021/10/19 23:43:55   Actual request no headers added: missing origin
2021/10/19 23:43:55 out GET /api/v1/items 235.621µs
```

Test the API endpoint with valid Authorization and Origin headers:

```
$ curl -D - -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dnZWRJbkFzIjoiYWRtaW4iLCJpYXQiOjE0MjI3Nzk2Mzh9.gzSraSYS8EXBxLN_oWnFSRgCzcmJmMjLiuyu5CSpyHI' -H 'Origin: https://my.dns.co' http://localhost:8080/api/v1/items
HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: https://my.dns.co
Content-Type: application/json
Server: MyApp-1.2.0
Vary: Origin
Date: Tue, 19 Oct 2021 21:45:00 GMT
Content-Length: 25

["item1","item2","item3"]
```

The corresponding server logs:

```
2021/10/19 23:45:00 in  GET /api/v1/items [::1]:54800
[cors] 2021/10/19 23:45:00 Handler: Actual request
[cors] 2021/10/19 23:45:00   Actual response added headers: map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[https://my.dns.co] Server:[MyApp-1.2.0] Vary:[Origin]]
2021/10/19 23:45:00 out GET /api/v1/items 341.432µs
```

## Low-level

See the [low-level example](examples/low-level/main.go).

The following code can be replaced by the high-level function `Server.RunServer()` presented in the previous chapter. The following code is intended to show that the Teal.Finance/Server can be customized to meet specific requirements.

```go
package main

import (
    "log"
    "net"
    "net/http"
    "time"

    "github.com/teal-finance/server"
    "github.com/teal-finance/server/chain"
    "github.com/teal-finance/server/cors"
    "github.com/teal-finance/server/export"
    "github.com/teal-finance/server/limiter"
    "github.com/teal-finance/server/opa"
    "github.com/teal-finance/server/reserr"
)

func main() {
    middlewares, connState := setMiddlewares()

    h := myHandler()
    h = middlewares.Then(h)

    runServer(h, connState)
}

func setMiddlewares() (middlewares chain.Chain, connState func(net.Conn, http.ConnState)) {
    // Uniformize error responses with API doc
    resErr := reserr.New("https://my.dns.co/doc")

    // Start a metrics server in background if export port > 0.
    // The metrics server is for use with Prometheus or another compatible monitoring tool.
    metrics := export.Metrics{}
    middlewares, connState = metrics.StartServer(9093, true)

    // Limit the input request rate per IP
    reqLimiter := limiter.New(10, 20, true, resErr)
    middlewares = middlewares.Append()

    // Endpoint authentication rules (Open Policy Agent)
    policy, err := opa.New(resErr, []string{"example-auth.rego"})
    if err != nil {
        log.Fatal(err)
    }

    // CORS
    allowedOrigins := []string{"https://my.dns.co"}

    middlewares = middlewares.Append(
        server.LogRequests,
        reqLimiter.Limit,
        server.Header("MyServerName-1.2.0"),
        policy.Auth,
        cors.Handle(allowedOrigins, true),
    )

    return middlewares, connState
}

// runServer runs in foreground the main server.
func runServer(h http.Handler, connState func(net.Conn, http.ConnState)) {
    server := http.Server{
        Addr:              ":8080",
        Handler:           h,
        TLSConfig:         nil,
        ReadTimeout:       1 * time.Second,
        ReadHeaderTimeout: 1 * time.Second,
        WriteTimeout:      1 * time.Second,
        IdleTimeout:       1 * time.Second,
        MaxHeaderBytes:    222,
        TLSNextProto:      nil,
        ConnState:         connState,
        ErrorLog:          log.Default(),
        BaseContext:       nil,
        ConnContext:       nil,
    }

    log.Print("Server listening on http://localhost", server.Addr)

    log.Fatal(server.ListenAndServe())
}
```
