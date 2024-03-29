# 🌈 ![Rainbow](doc/rainbow-chancery.png)

| ![logo](doc/small.png) | Dashboard for DeFi options trading highlighting market opportunities to ease strategies decision. With the Rainbow screener, users can scan CEX/DEX markets to fetch available options data. Empowered by the [Teal.Finance](https://teal.finance/) [team](https://teal.finance/#team). <br><br> [![Go Report Card](https://goreportcard.com/badge/github.com/teal-finance/rainbow)](https://goreportcard.com/report/github.com/teal-finance/rainbow) |
| ---------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

Rainbow is live at [**teal.finance/rainbow**](https://teal.finance/rainbow/)

Please share any feedback:

- GitHub [issue](https://github.com/teal-finance/rainbow/issues) or
- Email at Teal.Finance[at]pm.me or
- Twitter [@TealFinance](https://twitter.com/TealFinance).

## 🎯 Motivations

**Rainbow** is inspired by the following trends:

- Crypto-assets and DeFi becoming mainstream,
- Advent of the Internet of Blockchains’ world
  with more cross-chain applications and communications,
- Crypto options trading growing bigger
  than the spot market, like in traditional Finance.

The crypto derivatives markets are expected to grow a lot more in the upcoming years.
Specifically, the options markets will see the biggest growth,
because in Finance, Option market is much bigger than the underlying spot market.
This is lagging in Crypto, when we look at Deribit's volume (the main centralized venue for Crypto Options),
compared to Perpetuals future.

We participated to hackathons and have won super prices:
[Ethereum's EthGlobal Hackathon](https://showcase.ethglobal.com/ethonline2021/rainbow),
[Solana's Ignition Hackathon](https://devpost.com/software/rainbow-ai5p7m)
and [Encode Club's Hack DeFi](https://www.encode.club/hack-defi).

More info on our motivations:
[cryptonuage.com/posts/internet-of-decentralized-options](https://cryptonuage.com/posts/internet-of-decentralized-options/)

## 🏛️ Supported exchanges

- [Deribit](https://www.deribit.com):
  Centralized exchange.
  Main venue for crypto Options trading.

- [Delta Exchange](https://www.delta.exchange/):
  Centralized exchange for futures, options and interest
  rate swaps on 50 underlying cryptocurrencies.

- [Lyra](https://www.lyra.finance/) (Ethereum):
  DeFi protocol for trading options based on Layer 2,
  Optimism, using a custom AMM.

- [Thales](https://thalesmarket.io/) Optimism and Polygon,
  DeFi protocol for [binary options](https://wikiless.org/wiki/Binary_option) (exotic options).

- [Zeta](https://zeta.markets/) (Solana):
  DeFi options protocol built using Serum order books.

## 🏗 Deprecated exchanges

- [PsyOptions](https://www.psyoptions.io/) (Solana):
  DeFi options protocol build on Serum order books.
  But PsyOptions pivoted to DeFi Options Vaults (DOV).

- [Opyn](https://www.opyn.co/) (Ethereum):
  DeFi options protocol using TheGraph & 0x.
  But Opyn also pivoted to more vaults like products Squeeth, Crab...

## 🏎️ Run it from source code

Rainbow can be used as a [CLI](#🪃-run-the-cli)
or with a [frontend](#⚒️-makefile-for-server--web)
as our official instance:
<https://teal.finance/rainbow>.

### 📑 Requirements

- Go v1.22 (single requirement for the [CLI](#🪃-run-the-cli))
- Node v18
- Npm v9
- Docker v20 (optional)
- Podman v3 (optional)

Rainbow can be used in the [CLI](#🪃-run-the-cli),
without requiring Node, Npm, Docker…

[Snap](<https://en.wikipedia.org/wiki/Snap_(package_manager)>) provides a simple way to install these requirements on many Linux distributions:

    snap install go   --classic
    snap install node --classic
    go     version
    node --version

On Debian/Ubuntu, the command `sudo apt install golang` may install an older version.
List all available versions with:

    apt list --all-versions golang golang-1.*

To install a more recent Go version, you may try:

    sudo apt purge   golang*
    sudo apt install golang-1.22

However, you may not need to install Go and Node if you build/run Rainbow using the [Dockerfile](#🐋-container).

### 📥 Clone the Rainbow Git repo

    git clone https://github.com/teal-finance/rainbow
    cd rainbow

### 🪃 Run the CLI

The Rainbow project provides a Command Line Interface (CLI) to let users play from their terminal. The command `./cli` retrieves (few minutes) and prints a pretty nice table of all options.

    go run ./cmd/cli

![CLI screenshot](doc/cli.jpg)

### ⚒️ Makefile for server & web

See the `make help` [output](Makefile).

#### 🔧 Build

    make all     # Build both backend and frontend
    make server  # Build the backend only
    make front   # Build the frontend only
    make clean   # Clean all

#### 🚀 Run

The _run_ targets do not depend on the above _build_ targets.
You do not need to _make build_ before _make run_.

    make run        # Run the backend in dev mode
    make run-ui     # Run the frontend in dev mode

To serve the static website using the Rainbow backend only, as in production:

    make clean -j && make all -j && make run

#### 🐋 Container

    make container-run  # Build and run backend + frontend
    make container-rm   # Stop and remove all

The _container_ targets try to use `docker`, then `podman`.

The `container-run` also opens a browser window on <http://localhost:1111/>
and prints the container logs indefinitely.

Use CTRL+C to stop the log printing. This does not stop, the container.
To stop the container simply use `make container-rm`.

The two other _containers_ targets are not needed to be invoked manually:

    make container-build  # Build the container image
    make container-stop   # Stop/remove the container

The _container_ parameters can be customized: :slightly_smiling_face:

    make container-run expose=80 port=80 base=/rainbow/

### 🖐️ Manual build & run

#### Backend

The frontend requires the server API.

    go build ./cmd/server && ./server

    # or

    go run ./cmd/server     # same as "make run"

#### Frontend

To run the Vue3 frontend in dev mode.
Similar to `make run-ui`:

    cd frontend
    npm i
    npm run dev

In prod mode, the backend serves the web static files from `frontend/dist`.
Same as `make frontend/dist`:

    cd frontend
    npm i
    npm run build

Finally open <http://localhost:8090>

### 🐳 Manual build/run of the container

This Git repo provides a [Dockerfile](Dockerfile) for a secured and light all-in-one container image: 30 MB.

The image contains the hardened sever executable and the frontend static files.

The container enables by default the CORS, the export ports and a rate limiter.
Some of these features can be customized using environments variables.

The following commands configures the server listening on port 2222 (published to port 1111)
and frontend using CORS with `http://localhost:1111`.
See also the comments within the [Dockerfile](Dockerfile) for the details.

    podman build --build-arg addr=http://localhost:1111 --build-arg port=2222 -t rainbow .
    podman stop rainbow # if already running in background
    podman run --rm -p 0.0.0.0:1111:2222 -d --name rainbow rainbow
    podman logs --follow rainbow

Open <http://localhost:1111>.

The above manual commands can be obtained using the [Makefile](Makefile):

    make container-run addr=http://localhost:1111 expose=1111 port=2222

### 🚩 Server command line flags

Rainbow embeds a complete HTTP server,
including a rate limiter, an export port (Prometheus monitoring),
and more. For more details see the underlying project
[Teal.Finance/Garcon](https://github.com/teal-finance/garcon/).

    $ go build ./cmd/server
    $ ./server -h
    Usage of ./server:
      -addr string
          Schema and DNS used for doc URL and CORS, has precedence over MAIN_ADDR (default "http://localhost")
      -aes string
          128-bit AES key (32 hex digits) for the session cookies, has precedence over AES128
      -alert string
          Webhook endpoint to notify anomalies, has precedence over ALERT_URL
      -burst int
          Max requests during a burst, has precedence over REQ_BURST (default 22)
      -cex
          Enable the centralized exchanges: Deribit and Delta Exchange
      -dev
          Enable the developer mode (enabled by default if -addr and -port are not used)
      -dex
          Enable the decentralized exchanges: Lyra, Synquote and Zeta
      -exotic
          Enable the decentralized exchanges with binary options: Thales
      -exp int
          Export port for Prometheus, has precedence over EXP_PORT
      -form string
          Webhook endpoint to notify filled contact form, has precedence over WEBFORM_URL
      -hmac string
          HMAC-SHA256 key (64 hex digits) for the JWT tokens, has precedence over HMAC_SHA256
      -period duration
          Period to fetch market data from providers (default 10m0s)
      -port int
          API port, has precedence over MAIN_PORT (default 8090)
      -providers string
          Coma-separated list of providers, has precedence over PROVIDERS (default "ALL")
      -rate int
          Max requests per minute, has precedence over REQ_PER_MINUTE (default 88)
      -version
          Print version and exit
      -www string
          Folder of the web static files, has precedence over WWW_DIR (default "frontend/dist")

## 🛣️ API

The API is protected by a JWT.
You can reuse the anonymous JWT provided by the frontend.
To get the JWT, first visit the webpage
https://teal.finance/rainbow
then replace this former URL by
https://teal.finance/rainbow/v0/options

Moreover, you can also use the API by running the Rainbow backend in your machine:  

Using the [Dockerfile](#🐋-container):

    make container-run expose=8090

Using the Go compiler installed on your local machine:

    make run


### /v0/options

The [/v0/options](http://localhost:8090/v0/options) endpoint accepts optional parameters.

#### Filtering

Progressive filtering:

- [/v0/options/BTC](http://localhost:8090/v0/options/BTC)
  (only the BTC-based options)
- [/v0/options/BTC/2022-04-29](http://localhost:8090/v0/options/BTC/2022-04-29)
  (only the BTC-based options expiring on 2022-04-29)
- [/v0/options/BTC/2022-04-29/Deribit](http://localhost:8090/v0/options/BTC/2022-04-29/Deribit)
  (only the BTC-based options expiring on 2022-04-29 from Deribit)
- [/v0/options/BTC/2022-04-29/Deribit/csv](http://localhost:8090/v0/options/BTC/2022-04-29/Deribit/csv)
  (the same but in a [CSV](https://wikiless.org/wiki/Comma-separated_values "Comma-Separated Values") file,
  see the [supported formats](#supported-formats))

The general URL pattern:

    /v0/options/{asset}/{expiry}/{provider}/{format}

The API may be tested with [cURL](https://wikiless.org/wiki/cURL):

    curl localhost:8090/v0/options

    curl localhost:8090/v0/options/BTC

    curl localhost:8090/v0/options/BTC/2022-04-29

    curl localhost:8090/v0/options/BTC/2022-04-29/Deribit

    curl localhost:8090/v0/options/BTC/2022-04-29/Deribit/csv

#### Query string

Moreover, the parameters can also be passed using the query string:

- [/v0/options?asset=BTC&asset=ETH](http://localhost:8090/v0/options?asset=BTC&asset=ETH) (multiple underlying assets)
- [/v0/options?asset=BTC&expiry=2022&provider=Deribit&format=csv](http://localhost:8090/v0/options?asset=BTC&expiry=2022&provider=Deribit&format=csv)

#### POST method

The API also supports the POST [form](https://wikiless.org/wiki/Application/x-www-form-urlencoded#The_application/x-www-form-urlencoded_type "application/x-www-form-urlencoded"):

    curl localhost:8090/v0/options -d asset=BTC -d asset=ETH

    curl localhost:8090/v0/options -d asset=BTC -d expiry=2022 -d provider=Deribit -d format=csv

The POST [form](https://wikiless.org/wiki/Application/x-www-form-urlencoded#The_application/x-www-form-urlencoded_type "application/x-www-form-urlencoded")
is recommended for privacy reasons since it hides your query within the encrypted request body when using HTTPS.

#### Awesome expiry filtering

The expiry filtering matches for the beginning of the dates:

- [/v0/options/ALL](http://localhost:8090/v0/options/ALL) for `ALL` options and any expiry date
- [/v0/options/ALL/2022](http://localhost:8090/v0/options/ALL/2022) for expiry dates in 2022
- [/v0/options/ALL/2022-04](http://localhost:8090/v0/options/ALL/2022-04) for expiry dates in April 2022
- [/v0/options/ALL/2022-04-01](http://localhost:8090/v0/options/ALL/2022-04-01) for April 1st 2022

#### Supported formats

The `{format}` parameter is the last one at any position.

- [/v0/options](http://localhost:8090/v0/options) (default is JSON in the browser)
- [/v0/options/json](http://localhost:8090/v0/options/json) (download a file in JSON format)
- [/v0/options/csv](http://localhost:8090/v0/options/csv) (CSV file)
- [/v0/options/tsv](http://localhost:8090/v0/options/tsv)
  (TSV = [Tab-Separated Values](https://wikiless.org/wiki/Tab-separated_values))
- [/v0/options/BTC/csv](http://localhost:8090/v0/options/BTC/csv)
  (only the BTC-based options in CSV)
- [/v0/options/ALL/2022-06/csv](http://localhost:8090/v0/options/ALL/2022-06/csv)
  (all options expiring in June 2022 in CSV)

The current supported formats are JSON,
[CSV](https://wikiless.org/wiki/Comma-separated_values "Comma-Separated Values") and
[TSV](https://wikiless.org/wiki/Tab-separated_values "Tab-Separated Values").
Depending on user requests, more formats may be supported such as
[JSON-LD](https://github.com/piprate/json-gold "JSON for Linked Data"),
[JSON Lines](https://jsonlines.org/examples/ "JSONL") (or
[NDJSON](https://github.com/ndjson/ndjson.github.io/issues/1#issuecomment-109935996 "Newline-Delimited JSON"), see
[JSON streaming](https://wikiless.org/wiki/JSON%20streaming#Line-delimited_JSON)),
[Avro](https://github.com/linkedin/goavro),
[Parquet](https://github.com/xitongsys/parquet-go),
[DataFrame](https://github.com/rocketlaunchr/dataframe-go),
[Excel](https://github.com/qax-os/excelize),
[Google Sheets Docs](https://github.com/Iwark/spreadsheet).

#### HTTP header

The "Accept" HTTP header is also supported:

    curl localhost:8090/v0/options -H "Accept: application/json"

    curl localhost:8090/v0/options -H "Accept: text/csv"

    curl localhost:8090/v0/options -H "Accept: text/tsv"

#### Example response in JSON format

```js
{
  "name":     "ETH-2021-10-29 23:59:59-3200-PUT",
  "expiry":   "2021-10-29 23:59:59",
  "type":     "PUT",
  "asset":    "ETH",        // ETH, BTC, SOL
  "strike":   3200,
  "exchange": "DEX",        // Type: CEX or DEX
  "chain":    "Solana",     // Ethereum, Solana...
  "layer":    "L1",
  "provider": "PsyOptions", // Opyn, Lyra, Deribit...
  "currency": "USDC"        // Quote currency: BTC...
  "bid": [{
            "px": 13.3,     // Abbreviation for "price"
            "size": 5,      // Variant of "quantity"
          },
          {
            "px": 13.1,
            "size": 10,
          }],
  "ask": [{
            "px": 15.12,
            "size": 5,
          },
          {
            "px": 15.25,
            "size": 9,
          }]
}
```

#### Example response in CSV format

The CSV and TSV formats are limited to the two best order-books: Bid #1, Bid #2, Ask #1, Ask #2.

```csv
$ curl -s localhost:8090/v0/options/csv
Name,Expiry,Asset,Currency,Strike,Type,Provider,Layer,Chain,Bid Price #1,Bid Size #1,Bid Price #2,Bid Size #2,Ask Price #1,Ask Size #1,Ask Price #2,Ask Size #2
SOL-2022-04-01 08:00:00-120-Call,2022-04-01 08:00:00,SOL,USDC,120,CALL,Zeta,L1,Solana,0.9238,285.24,0.73,300,1.0463,285.24,1.1,300
SOL-2022-04-01 08:00:00-130-Call,2022-04-01 08:00:00,SOL,USDC,130,CALL,Zeta,L1,Solana,0.2187,285.24,0.09,300,0.3486,285.24,0.37,300
BTC-29APR22-40000-P,2022-04-29 08:00:00,BTC,BTC,40000,PUT,Deribit,–,–,736.51598,62.2,712.7574000000001,28.7,807.7917200000002,118.2,831.5503000000001,30.3
BTC-8APR22-44000-P,2022-04-08 08:00:00,BTC,BTC,44000,PUT,Deribit,–,–,617.72308,0.1,593.9645,49.8,641.48166,32.5,665.2402400000001,31.1
BTC-1APR22-42000-P,2022-04-01 08:00:00,BTC,BTC,42000,PUT,Deribit,–,–,71.27574000000001,0.1,47.517160000000004,75.4,95.03432000000001,82.7,118.79290000000002,50.9
```

### /v0/bff/cp

This endpoint has been inspired by the [BFF pattern](https://blog.bitsrc.io/e4fa965128bf)
(Backend for Frontend pattern) to simplify the frontend processing.
This endpoint should not used by API users because we may drop it in a further release.

```js
[ { "asset": "ETH",
    "expiry": "Dec 31",
    "provider": "Deribit",
    "call": { "bid": {"px": 1805, "size":  24},
              "ask": {"px": 1830, "size": 459}},
    "strike": 4400,
    "put":  { "bid": {"px": 1305, "size":  26},
              "ask": {"px": 1330, "size":  37}}
  },
  { "asset": "ETH",
    "expiry": "Dec 31",
    "provider": "Deribit",
    "call": { "bid": {"px": 1140, "size": 258},
              "ask": {"px": 1160, "size":  33}},
    "strike": 5200,
    "put":  { "bid": {"px": 1235, "size":  80},
              "ask": {"px":    0, "size":   0}}
  },
  // ...
  { "asset": "BTC",
    "expiry": "Nov 26",
    "provider": "Deribit",
    "call": { "bid": {"px": 2045, "size":  7.5},
              "ask": {"px": 2105, "size":  5.3},
    "strike": 50000,
    "put":  { "bid": {"px":  125, "size": 27.2},
              "ask": {"px":  135, "size": 70.2}}
} ]
```

### /v0/graphql

This endpoint is currently designed to meet the requirements of the frontend.
This endpoint should not yet used by API users because we may drop it in a further release.

However, other websites can display the Rainbow screener using this GraphQL API.
Please contact us. 😉

### /v0/graphiql (with an "i")

This interactive GraphQL endpoint is used in developer mode to tune the frontend GraphQL queries.

## 🗣️ Need something else?

Please share your thoughts, suggest ideas, request features,
contact us and join our efforts to build awesome tools:

- GitHub [issue](https://github.com/teal-finance/rainbow/issues)
- Teal.Finance[at]pm.me
- Twitter [@TealFinance](https://twitter.com/TealFinance)
