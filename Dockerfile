# Build:
#
#    DOCKER_BUILDKIT=1 
#    docker  build --build-arg url=http://my.dns.co:1111/rainbow/ -t rainbow .
#    podman  build --build-arg url=http://my.dns.co:1111/rainbow/ -t rainbow .
#    buildah build --build-arg url=http://my.dns.co:1111/rainbow/ -t rainbow .
#
# Run:
#
#    docker run -d --rm -p 0.0.0.0:1111:2222 -e EXP_PORT=9868 --name rainbow rainbow
#    podman run -d --rm -p 0.0.0.0:1111:2222 -e EXP_PORT=9868 --name rainbow rainbow

# --------------------------------------------------------------------
# Arguments to configure the build:
# - url is used:
#   - by frontend as --base (prefix) stripped by the reverse-proxy and as API URL
#   - by backend to set CORS origin ard Cookie domain
# - port is the backend listennig port for API and static website server

# Default values are for dev mode:
ARG url=http://localhost:8888/
ARG port=8888

# Example of Prod values:
# url  = https://my.dns.co/rainbow/
# port = 8888

# --------------------------------------------------------------------
FROM docker.io/node:16-alpine3.14 AS web_builder

WORKDIR /code

COPY frontend/package.json \
     frontend/yarn.lock   ./

RUN set -x                          &&\
    node --version                  &&\
    yarn --version                  &&\
    yarn install --frozen-lockfile  &&\
    yarn cache clean

COPY frontend/.eslintrc.json     \
     frontend/index.html         \
     frontend/postcss.config.js  \
     frontend/tailwind.config.js \
     frontend/tsconfig.json      \
     frontend/vite.config.ts     \
     frontend/.env              ./

COPY frontend/public public
COPY frontend/src    src

ARG url

RUN set -x                                                  &&\
    ls -lA                                                  &&\
    sed -e "s|^VITE_API_URL=.*|VITE_API_URL=$url|" -i .env  &&\
    head .env                                               &&\
    yarn build --base "$url"                                &&\
    yarn compress

# --------------------------------------------------------------------
FROM docker.io/golang:1.17-alpine3.14 AS go_builder

WORKDIR /code

COPY go.mod go.sum ./

RUN go version       &&\
    go mod download  &&\
    go mod verify

COPY cmd cmd
COPY pkg pkg

# Go build flags: https://shibumi.dev/posts/hardening-executables/
# "-s -w" removes all debug symbols: https://pkg.go.dev/cmd/link
RUN set -x                                                &&\
    ls -lA                                                &&\
    export GOOS=linux                                     &&\
    export CGO_ENABLED=0                                  &&\
    export GOFLAGS="-buildmode=pie -trimpath -modcacherw" &&\
    export GOLDFLAGS="-linkmode=external -s -w"           &&\
    go build ./cmd/server                                 &&\
    ls -sh server                                         &&\
    ldd server                                            &&\
    ./server -help  # smoke test

# To go further in Go hardening and FIPS 140-2 certification:
# https://www.linkedin.com/pulse/go-crypto-kubernetes-fips-140-2-fedramp-compliance-gokul-chandra
# https://github.com/golang/go/blob/dev.boringcrypto/README.boringcrypto.md
# https://hub.docker.com/r/goboring/golang/tags
# https://github.com/rancher/image-build-base/blob/master/Dockerfile.amd64


# --------------------------------------------------------------------
FROM docker.io/golang:1.17-alpine3.14 AS integrator

WORKDIR /target

# HTTPS root certificates (adds about 200 KB)
# Creaate user & group files
RUN set -x                                                  &&\
    mkdir -p                                 etc/ssl/certs  &&\
    cp -a /etc/ssl/certs/ca-certificates.crt etc/ssl/certs  &&\
    echo 'teal:x:5505:5505::/:' > etc/passwd                &&\
    echo 'teal:x:5505:'         > etc/group

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

ARG url
ARG port

COPY --chown=5505:5505 --from=integrator /target /

# Run as unprivileged
USER teal:teal

# Use UTC time zone by default
ENV TZ        UTC0
ENV WWW_DIR   /var/www
ENV MAIN_ADDR "$url"
ENV MAIN_PORT "$port"
ENV EXP_PORT  9868

EXPOSE "$port"
EXPOSE 9868

ENTRYPOINT ["/server"]
