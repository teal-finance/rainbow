# ![rainbow](images/rainbow-chancery.png)

&nbsp; | &nbsp;
---|---
![logo](images/small.png) | **Rainbow** is a dashboard for Decentralized Finance options trading. <br><br> It's developed during Solana's Ignition & Ethereum's EthGlobal Hackathons by some members of [Teal Finance](https://teal.finance/).

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

Most recently, PsyOptions (from the Solana ecosystem) has also been integrated. PsyOptions is the main DeFi options protocol.

More Ethereum and Solana protocols are planed to also be supported in the near future: Thales, Lyra, Hegic...

Now, using Rainbow, you can compare to arbitrage options across these markets, or simply get the best prices.

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

Next step is to implement on a proper UI that enables users to trade,
based on Typescript/Vue3.

## Build

    git clone https://github.com/teal-finance/rainbow
    cd rainbow
    go generate ./...
    go build ./cmd/rainbow/
    ./rainbow

![CLI screenshot](images/cli.jpg)

## Back-end API

To run Rainbow as a HTTP server, use command line argument `-port 8080`
or define the environment variable `API_PORT`.

Rainbow also provides an export port to be monitored by Prometheus
or other monitoring tools.

Rainbow also implements a configurable rate limiter.

    $ ./rainbow -help
    Usage of ./rainbow:
      -dev
          Run rainbow in dev. mode
      -exp int
          Export port for Prometheus, overseeds env. var. EXP_PORT
      -max-burst int
          Max requests during a burst, overseeds env. var. MAX_REQ_BURST (default 1)
      -max-rate int
          Max requests per minute, overseeds env. var. MAX_REQ_PER_MINUTE (default 6)
      -port int
          API port, overseeds env. var. API_PORT

See also the Dockerfile for a light container image (16.3 MB).
The container enables by default the API and export ports.
The Dockerfile supports Docker, but it is not compatible with Podman-3.3.1.

## Data structure

The API endpoint /v0/options replies an array of options.
Each option is described with the following JSON structure:

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
