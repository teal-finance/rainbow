module github.com/teal-finance/rainbow/poc

go 1.22

// Fix the Go-1.19 crash from an indirect dependency used by Solana libraries. See golang/go#54227
replace github.com/daaku/go.zipexe v1.0.1 => github.com/42wim/go.zipexe v0.0.0-20220806143830-85f957dec14b

// Organization name has changed within the repo URL
replace github.com/dfuse-io/logging v0.0.0-20210518215502-2d920b2ad1f2 => github.com/streamingfast/logging v0.0.0-20220511154537-ce373d264338

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/gagliardetto/binary v0.7.9
	github.com/gagliardetto/gofuzz v1.2.2
	github.com/gagliardetto/solana-go v1.8.4
	github.com/gagliardetto/treeout v0.1.4
	github.com/spewerspew/spew v0.0.0-20230513223542-89b69fbbe2bd
	github.com/streamingfast/solana-go v0.5.1-0.20220502224452-432fbe84aee8
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/ryanuber/columnize v2.1.2+incompatible
	github.com/teal-finance/rainbow v0.6.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/GeertJohan/go.rice v1.0.3 // indirect
	github.com/JohannesKaufmann/html-to-markdown v1.5.0 // indirect
	github.com/Khan/genqlient v0.7.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/PuerkitoBio/goquery v1.9.1 // indirect
	github.com/acmacalister/skittles v0.0.0-20160609003031-7423546701e1 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/carlmjohnson/flagx v0.22.2 // indirect
	github.com/carlmjohnson/versioninfo v0.22.5 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v0.7.0 // indirect
	github.com/cristalhq/base64 v0.1.2 // indirect
	github.com/daaku/go.zipexe v1.0.2 // indirect
	github.com/deckarep/golang-set/v2 v2.6.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/dfuse-io/logging v0.0.0-20210518215502-2d920b2ad1f2 // indirect
	github.com/ethereum/c-kzg-4844 v0.4.0 // indirect
	github.com/ethereum/go-ethereum v1.13.14 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/fgprof v0.9.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-chi/chi/v5 v5.0.12 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/google/pprof v0.0.0-20240227163752-401108e1b7e7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/rpc v1.2.1 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/mtraver/base91 v1.0.0 // indirect
	github.com/pkg/profile v1.7.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.50.0 // indirect
	github.com/prometheus/procfs v0.13.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/streamingfast/binary v0.0.0-20240116150400-b058b4376d53 // indirect
	github.com/streamingfast/logging v0.0.0-20230608130331-f22c91403091 // indirect
	github.com/supranational/blst v0.3.11 // indirect
	github.com/teal-finance/emo v0.0.0-20240312100810-1c5a4984c875 // indirect
	github.com/teal-finance/garcon v0.35.0 // indirect
	github.com/teal-finance/incorruptible v0.0.0-20240312103045-cd27cac96f11 // indirect
	github.com/teal-finance/quid v0.0.0-20230216061629-9c9b3ff5e887 // indirect
	github.com/tidwall/gjson v1.17.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tklauser/go-sysconf v0.3.13 // indirect
	github.com/tklauser/numcpus v0.7.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.15 // indirect
	github.com/ybbus/jsonrpc v2.1.2+incompatible // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.mongodb.org/mongo-driver v1.14.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/ratelimit v0.3.1 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/mod v0.16.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/term v0.18.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
