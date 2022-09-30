module github.com/teal-finance/rainbow/poc

go 1.19

// Fix the Go-1.19 crash from an indirect dependency used by Solana libraries. See golang/go#54227
replace github.com/daaku/go.zipexe v1.0.1 => github.com/42wim/go.zipexe v0.0.0-20220806143830-85f957dec14b

// Organization name has changed within the repo URL
replace github.com/dfuse-io/logging v0.0.0-20210518215502-2d920b2ad1f2 => github.com/streamingfast/logging v0.0.0-20220511154537-ce373d264338

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/gagliardetto/binary v0.7.1
	github.com/gagliardetto/gofuzz v1.2.2
	github.com/gagliardetto/solana-go v1.6.0
	github.com/gagliardetto/treeout v0.1.4
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/spewerspew/spew v0.0.0-20220201233537-1fb8bf5ed3d2
	github.com/streamingfast/solana-go v0.5.1-0.20220502224452-432fbe84aee8
	github.com/stretchr/testify v1.8.0
)

require github.com/ryanuber/columnize v0.0.0-20160712163229-9b3edd62028f

require (
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/GeertJohan/go.rice v1.0.2 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/aybabtme/rgbterm v0.0.0-20170906152045-cc83f3b3ce59 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/daaku/go.zipexe v1.0.1 // indirect
	github.com/dfuse-io/logging v0.0.0-20210518215502-2d920b2ad1f2 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/gorilla/rpc v1.2.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/onsi/gomega v1.17.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/streamingfast/binary v0.0.0-20210928223119-44fc44e4a0b5 // indirect
	github.com/streamingfast/logging v0.0.0-20220813175024-b4fbb0e893df // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/ybbus/jsonrpc v2.1.2+incompatible // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be // indirect
	golang.org/x/net v0.0.0-20220927171203-f486391704dc // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/term v0.0.0-20220919170432-7a66f970e087 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
