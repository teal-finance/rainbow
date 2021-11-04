# ![rainbow](doc/rainbow-chancery.png)

![logo](doc/small.png) | Rainbow is a dashboard for Decentralized Finance options trading. <br><br> It's developed during Solana's Ignition & Ethereum's EthGlobal Hackathons by members of [Teal.Finance](https://teal.finance/).
---|---

## Motivation

The crypto derivatives markets are expected
to grow a lot more in the upcoming years.
Specifically, the options markets will see the bigger growth.

In traditional finance, option market is much bigger
than the underlying assets market.
This is lagging in Crypto, when we look at Deribit's volume,
the main (centralized) venue for Crypto Options,
compared to Perpetuals future.

**Rainbow** profits from the following trends:

* Crypto-assets and DeFi becoming mainstream,

* DeFi composability & complexity abstraction
  for easy onboarding of users,

* Advent of the Internet of Blockchains’ world
  with more cross-chain applications and communications,

* Crypto options trading growing (hopefully) bigger
  than the spot market, like in tradFi.

## Target

**Rainbow** is a first step in that direction:
a place where users, traders and market makers can see
the options, prices, expiries, liquidities across layers L1/L2.

## Current status

The current version is a tool to compare market data
across multiple venues, CEX and DEXes.

Deribit is supported since the beginning, because Deribit
is the main options trading place.

Opyn is also supported, because we are active users of this protocol.
Opyn use the 0x Protocol for the exchange of their options.
Rainbow uses both Opyn and 0x APIs to retrieve their trading data.
TheGraph is also used to list the available options from Opyn.

Most recently, PsyOptions (from the Solana ecosystem) has also been integrated.
PsyOptions is the main DeFi options protocol.

More Ethereum and Solana protocols are planed to also be supported
in the near future: Thales, Lyra, Hegic...

Now, using Rainbow, you can compare to arbitrage options
across these markets, or simply get the best prices.

## Technology

The back-end is developed in Go.

Deribit API is well documented, and the data retrieval
was pretty straightforward to implement.
Deribit even have an API playground, which all projects should also provide.

To support Opyn, Rainbow retrieves the options list using the TheGraph API.
This is our first GraphQL client implementation in Go and we spent days
to test and compare the different solutions:
we are proud to use the GraphQL state-of-the-art in Go,
based on the library <https://github.com/Khan/genqlient>
with type-safe code generation.
We got help from the Opyn team for the query examples.

Opyn support also requires to use the 0x protocol.
We have battled to correctly get the bid/ask prices from 0x API.
We got some help from the 0x team to identify which API call was the most suitable.
This help advances our work a lot and enables us to add,
in a near future, other protocols like Thales.

To support PsyOptions, we took our first plunge
in the Solana ecosystem to understand and use the Serum Go library.

See also the Go documentation: <https://pkg.go.dev/github.com/teal-finance/rainbow>

Next step is to implement on a proper UI that enables users to trade,
based on Typescript/Vue3.

## Build

    git clone https://github.com/teal-finance/rainbow
    cd rainbow
    go generate ./...

### CLI

You may just use the pretty nice table printed by the command `./cli`.

    go build ./cmd/cli && ./cli

![CLI screenshot](doc/cli.jpg)

### Back-end

If you prefer the API, use `./server`.

    go build ./cmd/server && ./server -dev

The flag `-dev` enables CORS for `http://localhost:*`.

### Front-end

This project also provides a pretty nice Vue3 front-end that uses the API.

    cd frontend
    yarn
    yarn dev

## Container

See also the [Dockerfile](Dockerfile) for a light container image: 30 MB.

The image contains the hardened sever executable (with dynamic library) and the front-end.

The container enables by default the CORS, the export ports and a rate limiter.
Some of these features can be customized using environments variables.

The Dockerfile has been successfully tested with Docker-20.10.8 and Podman-3.3.1.

## Command line flags

Rainbow provides a complete HTTP server,
including a rate limiter, an export port (Prometheus monitoring),
and more. For more details see the underlying project
[Teal.Finance/Server](https://github.com/teal-finance/teal/).

```
$ go build ./cmd/server
$ ./server -help
Usage of ./server:
  -addr string
        Schema and DNS used for doc URL and CORS, has precedence over MAIN_ADDR (default "http://localhost")
  -burst int
        Max requests during a burst, has precedence over REQ_BURST (default 10)
  -dev
        Run rainbow in dev. mode
  -exp int
        Export port for Prometheus, has precedence over EXP_PORT
  -opa string
        Policy files (comma-separated filenames) for the Open Policy Agent using the Datalog/Rego format
  -port int
        API port, has precedence over MAIN_PORT (default 8090)
  -rate int
        Max requests per minute, has precedence over REQ_PER_MINUTE (default 30)
  -www string
        Folder of the web static files, has precedence over WWW_DIR (default "frontend/dist")
```

## API

The API endpoints are designed to simplify the front-end source code.

![Idea of table to display on the web front-end](doc/frontend-table.png)

### /v0/expiries

<http://localhost:8090/v0/expiries>

List the available values for the two selectors above the table: Expiration date and Asset.

```json
{"expiries":
{"BTC":{"2021-10-29 23:59:59":16,"2021-11-26 08:00:00":16,"2021-12-31 08:00:00":12},
"ETH":{"2021-10-29 23:59:59":11,"2021-11-26 08:00:00":16,"2021-12-31 08:00:00":16},
"WBTC":{"2021-11-05 08:00:00":2,"2021-11-26 08:00:00":1},
"WETH":{"2021-11-04 08:00:00":1,"2021-11-05 08:00:00":2}}}
```

### /v0/tables/default

<http://localhost:8090/v0/tables/default>

Display the default table with the default selector values. The response also provide these selector values:

```json
{"asset":"BTC","expiry":"2021-10-29 23:59:59","table":
[{"p":"PsyOptions","cb":128.9,"cp":0,"ca":134.09,"cc":0,"cv":153,"co":0,"s":50000,"pb":7.51,"pp":0,"pa":8.63,"pc":0,"pv":120,"po":0},
{"p":"PsyOptions","cb":13.32,"cp":0,"ca":15.07,"cc":0,"cv":158,"co":0,"s":100000,"pb":0,"pp":0,"pa":0,"pc":0,"pv":0,"po":0},
{"p":"PsyOptions","cb":221.24,"cp":0,"ca":228.32,"cc":0,"cv":162,"co":0,"s":40000,"pb":1,"pp":0,"pa":1.21,"pc":0,"pv":57,"po":0},
{"p":"PsyOptions","cb":19.94,"cp":0,"ca":22.22,"cc":0,"cv":143,"co":0,"s":70000,"pb":95.64,"pp":0,"pa":99.8,"pc":0,"pv":149,"po":0},
{"p":"PsyOptions","cb":11.68,"cp":0,"ca":13.37,"cc":0,"cv":113,"co":0,"s":75000,"pb":136.58,"pp":0,"pa":141.75,"pc":0,"pv":125,"po":0},
{"p":"PsyOptions","cb":34.21,"cp":0,"ca":37.37,"cc":0,"cv":137,"co":0,"s":65000,"pb":60.54,"pp":0,"pa":64.18,"pc":0,"pv":114,"po":0},
{"p":"PsyOptions","cb":56.91,"cp":0,"ca":60.3,"cc":0,"cv":126,"co":0,"s":60000,"pb":33.87,"pp":0,"pa":36.74,"pc":0,"pv":145,"po":0},
{"p":"PsyOptions","cb":89.1,"cp":0,"ca":93.18,"cc":0,"cv":151,"co":0,"s":55000,"pb":16.71,"pp":0,"pa":18.71,"pc":0,"pv":139,"po":0}]}
```

### /v0/tables/{asset}/{expiry}

[http://localhost:8090/v0/tables/BTC/2021-11-26 08:00:00](http://localhost:8090/v0/tables/BTC/2021-11-26%2008:00:00)

Display the table when the user plays with the selectors.

The response has the same structure as the previous endpoint.

### /v0/options

<http://localhost:8090/v0/options>

List all the options and their order books.

```js
{
  "Name": "ETH-2021-10-29 23:59:59-3200-PUT",
  "Expiry": "2021-10-29 23:59:59",
  "Type": "PUT",
  "Asset": "ETH",             // ETH, BTC, SOL
  "Strike": 3200,
  "ExchangeType": "DEX",      // CEX or DEX
  "Chain": "Solana",          // Ethereum, Solana...
  "Layer": "L1",
  "Provider": "PsyOptions",   // Opyn, Lyra, Thales, Deribit, Psyoptions
  "Bid": [
    {
      "Price": 13.3,
      "Quantity": 5,
      "QuoteCurrency": "USDC" // ETH, BTC...
    },
    {
      "Price": 13.1,
      "Quantity": 10,
      "QuoteCurrency": "USDC"
    }
  ],
  "Ask": [
    {
      "Price": 15.12,
      "Quantity": 5,
      "QuoteCurrency": "USDC"
    },
    {
      "Price": 15.25,
      "Quantity": 9,
      "QuoteCurrency": "USDC"
    }
  ]
}
```
