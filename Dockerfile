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
FROM docker.io/node:18-alpine AS web_builder

WORKDIR /code

COPY frontend/package.json \
     frontend/yarn.lock   ./

RUN set -ex                         ;\
    node --version                  ;\
    yarn versions                   ;\
    yarn install --frozen-lockfile  ;\
    yarn cache clean

COPY frontend/index.html         \
     frontend/postcss.config.js  \
     frontend/tailwind.config.js \
     frontend/tsconfig.json      \
     frontend/vite.config.ts     \
     frontend/.env              ./

COPY frontend/public public
COPY frontend/src    src

ARG addr
ARG base

# No cache folder "/.gzipper" enabled by "gzipper compress --incremental"
ENV GZIPPER_INCREMENTAL     0
ENV GZIPPER_VERBOSE         0
ENV GZIPPER_SKIP_COMPRESSED 1

RUN set -ex                                            ;\
    ls -lA                                             ;\
    sed -e "s|^VITE_ADDR=.*|VITE_ADDR=$addr|" -i .env  ;\
    head .env                                          ;\
    yarn build --base "$base"                          ;\
    yarn compress

# --------------------------------------------------------------------
FROM docker.io/golang:1.18-alpine AS go_builder

WORKDIR /code

COPY go.mod go.sum ./

RUN set -ex          ;\
    go version       ;\
    go mod download  ;\
    go mod verify

COPY cmd cmd
COPY pkg pkg

# Go build flags: https://shibumi.dev/posts/hardening-executables/
# "-s -w" removes all debug symbols: https://pkg.go.dev/cmd/link
RUN set -ex                                               ;\
    ls -lA                                                ;\
    export GOOS=linux                                     ;\
    export CGO_ENABLED=0                                  ;\
    export GOFLAGS="-buildmode=pie -trimpath -modcacherw" ;\
    export GOLDFLAGS="-linkmode=external -s -w"           ;\
    go build ./cmd/server                                 ;\
    ls -sh server                                         ;\
    ldd server                                            ;\
    ./server -help  # smoke test

# To go further in Go hardening and FIPS 140-2 certification:
# https://www.linkedin.com/pulse/go-crypto-kubernetes-fips-140-2-fedramp-compliance-gokul-chandra
# https://github.com/golang/go/blob/dev.boringcrypto/README.boringcrypto.md
# https://hub.docker.com/r/goboring/golang/tags
# https://github.com/rancher/image-build-base/blob/master/Dockerfile.amd64


# --------------------------------------------------------------------
FROM docker.io/golang:1.18-alpine AS integrator

WORKDIR /target

# HTTPS root certificates (adds about 200 KB)
# Create user & group files
RUN set -ex                                                 ;\
    mkdir -p                                 etc/ssl/certs  ;\
    cp -a /etc/ssl/certs/ca-certificates.crt etc/ssl/certs  ;\
    echo "teal:x:$uid:$uid::/:" > etc/passwd                ;\
    echo "teal:x:$uid:"         > etc/group

# Static website and back-end
COPY --from=web_builder /code/dist   var/www
COPY --from=go_builder  /code/server .

# Copy possible dynamic libs because we use CGO_ENABLED=1
# https://stackoverflow.com/q/62817082
RUN mkdir lib        &&\
    ldd server        |\
    while read f rest ;\
    do cp -v "$f" lib ;\
    done             &&\
    ls -lA

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
