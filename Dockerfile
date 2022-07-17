# Build:
#
#    export DOCKER_BUILDKIT=1
#    docker  build --build-arg addr=http://my.dns.co:1111 --build-arg port=2222 --build-arg base=/rainbow/ -t rainbow .
#    podman  build --build-arg addr=http://my.dns.co:1111 --build-arg port=2222 --build-arg base=/rainbow/ -t rainbow .
#    buildah build --build-arg addr=http://my.dns.co:1111 --build-arg port=2222 --build-arg base=/rainbow/ -t rainbow .
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
FROM docker.io/golang:1.18 AS version

WORKDIR /code

COPY go.mod go.sum ./

RUN set -ex          ;\
    go version       ;\
    go mod download  ;\
    go mod verify

COPY cmd cmd
COPY pkg pkg
COPY .git .git

# Go build flags: https://shibumi.dev/posts/hardening-executables/
# "-s -w" removes all debug symbols: https://pkg.go.dev/cmd/link
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
    go mod download  ;\
    go mod verify

COPY cmd cmd
COPY pkg pkg
COPY .git .git

# Go build flags: https://shibumi.dev/posts/hardening-executables/
# "-s -w" removes all debug symbols: https://pkg.go.dev/cmd/link
RUN set -ex                                                                      ;\
    v="$(cat version.txt)"                                                       ;\
    export CGO_ENABLED=0                                                         ;\
    export GOFLAGS="-trimpath -modcacherw"                                       ;\
    export GOLDFLAGS="-d -s -w -extldflags=-static"                              ;\
    go build -v -ldflags="-X 'github.com/teal-finance/garcon.V=$v'" ./cmd/server ;\
    ls -sh server                                                                ;\
    ./server -version  # smoke test

# To enable Go hardening (FIPS 140-2 certification) set:
# GOFLAGS="-buildmode=pie -trimpath -modcacherw"
# GOLDFLAGS="-linkmode=external -s -w"
# https://www.linkedin.com/pulse/go-crypto-kubernetes-fips-140-2-fedramp-compliance-gokul-chandra
# https://github.com/golang/go/blob/dev.boringcrypto/README.boringcrypto.md
# https://hub.docker.com/r/goboring/golang/tags
# https://github.com/rancher/image-build-base/blob/master/Dockerfile.amd64

# --------------------------------------------------------------------
FROM docker.io/node:18-alpine AS web_builder

WORKDIR /code

COPY frontend/package.json \
    frontend/yarn.lock    ./

RUN set -ex                         ;\
    node --version                  ;\
    yarn versions                   ;\
    yarn install --frozen-lockfile  ;\
    yarn cache clean

COPY frontend/index.html        \
    frontend/postcss.config.js  \
    frontend/tailwind.config.js \
    frontend/tsconfig.json      \
    frontend/vite.config.ts     \
    frontend/.env              ./

COPY frontend/public public
COPY frontend/src    src

COPY --from=version /code/version.txt ./

ARG addr
ARG base

# No cache folder "/.gzipper" enabled by "gzipper compress --incremental"
ENV GZIPPER_INCREMENTAL     0
ENV GZIPPER_VERBOSE         0
ENV GZIPPER_SKIP_COMPRESSED 1

RUN set -ex                                            ;\
    ls -lA                                             ;\
    v="$(cat version.txt)"                             ;\
    sed -e "s|^VITE_VERS=.*|VITE_VERS=$v|"    -i .env  ;\
    sed -e "s|^VITE_ADDR=.*|VITE_ADDR=$addr|" -i .env  ;\
    sed -e "s|^VITE_BASE=.*|VITE_BASE=$base|" -i .env  ;\
    head .env                                          ;\
    yarn build --base "$base"                          ;\
    yarn compress

# --------------------------------------------------------------------
FROM docker.io/golang:1.18 AS integrator

WORKDIR /target

# Copy HTTPS root certificates (200 KB) + Create user/group files 
RUN set -ex                                                 ;\
    mkdir -p                                 etc/ssl/certs  ;\
    cp -a /etc/ssl/certs/ca-certificates.crt etc/ssl/certs  ;\
    echo "teal:x:$uid:$uid::/:" > etc/passwd                ;\
    echo "teal:x:$uid:"         > etc/group

# Copy static website and back-end executable
COPY --from=web_builder /code/dist   var/www
COPY --from=go_builder  /code/server .

# The following commented code copies the dynamic libs
# when Go build uses CGO_ENABLED=1 https://stackoverflow.com/q/62817082
# RUN mkdir lib           &&\
#     ldd server           |\
#     while read f rest    ;\
#     do cp -v "$f" lib    ;\
#     done                &&\
#     ls -lA

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
ENV TZ        UTC0
ENV WWW_DIR   /var/www
ENV MAIN_ADDR "$addr$base"
ENV MAIN_PORT "$port"
ENV EXP_PORT  9868
ENV HMAC_SHA256 C1C2C3C4C5C6C7C8C9C0CACBCCCDCECFC1C2C3C4C5C6C7C8C9C0CACBCCCDCECF

EXPOSE "$port"
EXPOSE 9868

ENTRYPOINT ["/server"]
