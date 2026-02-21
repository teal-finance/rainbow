# Build:
#
#    docker  build --build-arg addr=http://my.dns.co:1111 --build-arg port=2222 --build-arg base=/rainbow/ -f Containerfile -t rainbow .
#    podman  build --build-arg addr=http://my.dns.co:1111 --build-arg port=2222 --build-arg base=/rainbow/ -f Containerfile -t rainbow .
#    buildah build --build-arg addr=http://my.dns.co:1111 --build-arg port=2222 --build-arg base=/rainbow/ -f Containerfile -t rainbow .
#
# Run:
#
#    docker run -d --rm -p 0.0.0.0:1111:2222 -e EXP_PORT=9868 --name rainbow rainbow
#    podman run -d --rm -p 0.0.0.0:1111:2222 -e EXP_PORT=9868 --name rainbow rainbow

# --------------------------------------------------------------------
# Arguments:
# - addr is used:
#   - by frontend as the API scheme and host.
#   - by backend to set CORS origin ard Cookie domain.
# - base is passed to frontend with --base (the path prefix stripped by the reverse-proxy).
#   The frontend also uses it to compose the API URL.
# - port is used by backend as the listening port to serve API and static website files.

# Default values are for dev mode:
ARG addr=http://localhost:8888
ARG base=/
ARG port=8888
ARG uid=5505

# Example of some Prod values:
# addr = https://my.dns.co
# base = /rainbow/

# --------------------------------------------------------------------
FROM docker.io/golang:1.26 AS version

WORKDIR /code

COPY .git .git

RUN set -ex                                           ;\
    t="$(git describe --tags --abbrev=0 --always)"    ;\
    b="$(git branch --show-current)"                  ;\
    [ _$b = _main ] && b="" || b="-$b"                ;\
    n="$(git rev-list --count "$t"..)"                ;\
    [ "$n" -eq 0 ] && n="" || n="+$n"                 ;\
    v="$t$b$n"                                        ;\
    echo "Compute version string: $v"                 ;\
    echo -n "$v" > version.txt

# --------------------------------------------------------------------
FROM version AS go_builder

WORKDIR /code

COPY go.mod go.sum ./

RUN set -ex          ;\
    go version       ;\
    go mod download

COPY cmd cmd
COPY pkg pkg

# --mount=type -> https://docs.docker.com/build/cache/optimize/#use-cache-mounts
# Go build flags: "-s -w" removes all debug symbols: https://pkg.go.dev/cmd/link
#                https://shibumi.dev/posts/hardening-executables/
# GOAMD64=v3 --> https://go.dev/wiki/MinimumRequirements#amd64
RUN --mount=type=cache,target=/go/pkg/mod                \
    --mount=type=cache,target=/root/.cache/go-build      \
    set -ex                                             ;\
    ls -lShA                                            ;\
    case "$(grep flags -m1 /proc/cpuinfo)" in           ;\
        *" avx512f "*)  export GOAMD64=v4;;             ;\
        *" avx2 "*)     export GOAMD64=v3;;             ;\
        *" sse2 "*)     export GOAMD64=v2;;             ;\
    esac                                                ;\
    v="$(cat version.txt)"                              ;\
    flags="-X 'github.com/lynxai-team/garcon/vv.V=$v' -linkmode=internal"  ;\
    CGO_ENABLED=0 GOFLAGS="-trimpath -modcacherw"        \
    GOLDFLAGS="-d -s -w -extldflags=-static"             \
    go build -a -v -ldflags="$flags" ./cmd/server       ;\
    ls -lh server                                       ;\
    ./server -version  # smoke test

# To enable Go hardening (FIPS 140-2 certification) set:
# GOFLAGS="-buildmode=pie -trimpath -modcacherw"
# GOLDFLAGS="-linkmode=external -s -w"
# https://www.linkedin.com/pulse/go-crypto-kubernetes-fips-140-2-fedramp-compliance-gokul-chandra
# https://github.com/golang/go/blob/dev.boringcrypto/README.boringcrypto.md
# https://hub.docker.com/r/goboring/golang/tags
# https://github.com/rancher/image-build-base/blob/master/Dockerfile.amd64

# --------------------------------------------------------------------
# https://hub.docker.com/_/node
FROM docker.io/node:24-alpine AS web_builder

WORKDIR /code

COPY frontend/package.json       \
     frontend/package-lock.json ./

RUN set -ex                     ;\
    node --version              ;\
    npm --version               ;\
    npm cache verify            ;\
    npm ci --prefer-online --no-audit --no-fund
    # ci = clean-install => read-only package-lock.json
    # --prefer-offline   => use cache first, else download for internet
    # --prefer-online    => download for internet, except 
    # --no-fund          => no funding display
    # --no-audit         => no security audit
    # --only=production  => exclude devDependencies
    # --omit=dev         => do not install dev tools (tsc?)

COPY frontend/index.html         \
     frontend/postcss.config.js  \
     frontend/tailwind.config.js \
     frontend/tsconfig.json      \
     frontend/vite.config.mts    \
     frontend/.env              ./

COPY frontend/public public
COPY frontend/src    src

COPY --from=version /code/version.txt ./

ARG addr
ARG base

# No cache folder "/.gzipper" enabled by "gzipper compress --incremental"
ENV GZIPPER_INCREMENTAL=0
ENV GZIPPER_VERBOSE=0
ENV GZIPPER_SKIP_COMPRESSED=1

RUN set -ex                                            ;\
    ls -lShA                                           ;\
    v="$(cat version.txt)"                             ;\
    sed -e "s|^VITE_VERS=.*|VITE_VERS=$v|"    -i .env  ;\
    sed -e "s|^VITE_ADDR=.*|VITE_ADDR=$addr|" -i .env  ;\
    sed -e "s|^VITE_BASE=.*|VITE_BASE=$base|" -i .env  ;\
    head .env                                          ;\
    npm run build                                      ;\
    npm run compress

# --------------------------------------------------------------------
FROM docker.io/golang:1.26 AS integrator

WORKDIR /target

# Copy HTTPS root certificates (200 KB) + Create user/group files
RUN set -ex                                                 ;\
    mkdir -p                                 etc/ssl/certs  ;\
    cp -a /etc/ssl/certs/ca-certificates.crt etc/ssl/certs  ;\
    echo "teal:x:$uid:$uid::/:" > etc/passwd                ;\
    echo "teal:x:$uid:"         > etc/group

# Copy static website and backend executable
COPY --from=web_builder /code/dist   var/www
COPY --from=go_builder  /code/server .

# --------------------------------------------------------------------
FROM scratch AS final

ARG addr
ARG base
ARG port
ARG uid

COPY --chown=$uid:$uid --from=integrator /target /

# Run as unprivileged
USER $uid:$uid

# Use UTC time zone by default
ENV TZ=UTC0
ENV WWW_DIR=/var/www
ENV MAIN_ADDR="$addr$base"
ENV MAIN_PORT="$port"
ENV EXP_PORT=9868
ENV HMAC_SHA256=C1C2C3C4C5C6C7C8C9C0CACBCCCDCECFC1C2C3C4C5C6C7C8C9C0CACBCCCDCECF

EXPOSE "$port"
EXPOSE 9868

ENTRYPOINT ["/server"]
