module github.com/teal-finance/rainbow

go 1.21

// Fix the Go-1.19 crash from an indirect dependency used by Solana libraries. See golang/go#54227
replace github.com/daaku/go.zipexe v1.0.1 => github.com/42wim/go.zipexe v0.0.0-20220806143830-85f957dec14b

// Organization name has changed within the repo URL
replace github.com/dfuse-io/logging v0.0.0-20221209193439-bff11742bf4c => github.com/streamingfast/logging v0.0.0-20220511154537-ce373d264338

require (
	github.com/Khan/genqlient v0.5.0
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.12.1
	github.com/friendsofgo/graphiql v0.2.2
	github.com/gagliardetto/binary v0.7.8
	github.com/gagliardetto/gofuzz v1.2.2
	github.com/gagliardetto/solana-go v1.8.2
	github.com/gagliardetto/treeout v0.1.4
	github.com/go-chi/chi/v5 v5.0.8
	github.com/google/go-cmp v0.5.9
	github.com/gookit/color v1.5.3
	github.com/graphql-go/graphql v0.8.1
	github.com/graphql-go/handler v0.2.3
	github.com/jedib0t/go-pretty/v6 v6.4.6
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de
	github.com/spewerspew/spew v0.0.0-20220201233537-1fb8bf5ed3d2
	github.com/streamingfast/solana-go v0.5.1-0.20220502224452-432fbe84aee8
	github.com/stretchr/testify v1.8.2
	github.com/teal-finance/emo v0.0.0-20221020231838-605890cfa1bd
	github.com/teal-finance/garcon v0.33.0
)

require (
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/GeertJohan/go.rice v1.0.3 // indirect
	github.com/JohannesKaufmann/html-to-markdown v1.3.7 // indirect
	github.com/PuerkitoBio/goquery v1.8.1 // indirect
	github.com/acmacalister/skittles v0.0.0-20160609003031-7423546701e1 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/alexflint/go-arg v1.4.3 // indirect
	github.com/alexflint/go-scalar v1.2.0 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/aybabtme/rgbterm v0.0.0-20170906152045-cc83f3b3ce59 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/carlmjohnson/flagx v0.22.2 // indirect
	github.com/carlmjohnson/versioninfo v0.22.4 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cristalhq/base64 v0.1.2 // indirect
	github.com/daaku/go.zipexe v1.0.2 // indirect
	github.com/deckarep/golang-set/v2 v2.3.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/dfuse-io/logging v0.0.0-20221209193439-bff11742bf4c // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/felixge/fgprof v0.9.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/pprof v0.0.0-20230406165453-00490a63f317 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/rpc v1.2.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/holiman/uint256 v1.2.3 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/mtraver/base91 v1.0.0 // indirect
	github.com/pkg/profile v1.7.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.15.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/rs/cors v1.9.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/streamingfast/binary v0.0.0-20230125145418-8ef275399723 // indirect
	github.com/streamingfast/logging v0.0.0-20221209193439-bff11742bf4c // indirect
	github.com/teal-finance/incorruptible v0.0.0-20221218151336-b7408a3dd11b // indirect
	github.com/teal-finance/quid v0.0.0-20230216061629-9c9b3ff5e887 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.1 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	github.com/ybbus/jsonrpc v2.1.2+incompatible // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.mongodb.org/mongo-driver v1.11.4 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/term v0.8.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.9.1 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
