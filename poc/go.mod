module github.com/teal-finance/rainbow/poc

go 1.19

// Fix the Go-1.19 crash from an indirect dependency used by Solana libraries. See golang/go#54227
replace github.com/daaku/go.zipexe v1.0.1 => github.com/42wim/go.zipexe v0.0.0-20220806143830-85f957dec14b

require (
	github.com/Khan/genqlient v0.5.0
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.10.25
	github.com/friendsofgo/graphiql v0.2.2
	github.com/gagliardetto/binary v0.7.1
	github.com/gagliardetto/gofuzz v1.2.2
	github.com/gagliardetto/solana-go v1.6.0
	github.com/gagliardetto/treeout v0.1.4
	github.com/go-chi/chi/v5 v5.0.7
	github.com/google/go-cmp v0.5.9
	github.com/gookit/color v1.5.2
	github.com/graphql-go/graphql v0.8.0
	github.com/graphql-go/handler v0.2.3
	github.com/jedib0t/go-pretty/v6 v6.3.9
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de
	github.com/spewerspew/spew v0.0.0-20220201233537-1fb8bf5ed3d2
	github.com/streamingfast/solana-go v0.5.1-0.20220502224452-432fbe84aee8
	github.com/stretchr/testify v1.8.0
	github.com/teal-finance/emo v0.0.0-20220917213604-60c2a1947b7d
	github.com/teal-finance/garcon v0.30.0
)

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.13.14 // indirect
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/GeertJohan/go.rice v1.0.2 // indirect
	github.com/JohannesKaufmann/html-to-markdown v1.3.6 // indirect
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/acmacalister/skittles v0.0.0-20160609003031-7423546701e1 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/alexflint/go-arg v1.4.3 // indirect
	github.com/alexflint/go-scalar v1.1.0 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/aybabtme/rgbterm v0.0.0-20170906152045-cc83f3b3ce59 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.1 // indirect
	github.com/carlmjohnson/flagx v0.22.2 // indirect
	github.com/carlmjohnson/versioninfo v0.22.4 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cristalhq/base64 v0.1.2 // indirect
	github.com/daaku/go.zipexe v1.0.1 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/dfuse-io/logging v0.0.0-20210518215502-2d920b2ad1f2 // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/rpc v1.2.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/mtraver/base91 v1.0.0 // indirect
	github.com/onsi/gomega v1.20.2 // indirect
	github.com/pkg/profile v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/ryanuber/columnize v2.1.2+incompatible // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/streamingfast/binary v0.0.0-20210928223119-44fc44e4a0b5 // indirect
	github.com/streamingfast/logging v0.0.0-20220813175024-b4fbb0e893df // indirect
	github.com/teal-finance/incorruptible v0.0.0-20220924215045-e44628c14a6f // indirect
	github.com/teal-finance/quid v0.0.0-20220916181304-905aa35a284f // indirect
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569 // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.5.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.1 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	github.com/ybbus/jsonrpc v2.1.2+incompatible // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.2.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220927171203-f486391704dc // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/term v0.0.0-20220919170432-7a66f970e087 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
