# Build:
#
#    DOCKER_BUILDKIT=1 
#    docker  build --build-arg addr=http://myapp.co --build-arg port=8088 -t rainbow .
#    podman  build --build-arg addr=http://myapp.co --build-arg port=8088 -t rainbow .
#    buildah build --build-arg addr=http://myapp.co --build-arg port=8088 -t rainbow .
#
# Run:
#
#    docker run -p 0.0.0.0:8088:8088 -d -e EXP_PORT=9868 --name rainbow rainbow
#    podman run -p 0.0.0.0:8088:8088 -d -e EXP_PORT=9868 --name rainbow rainbow
#
# One command line:
#
#    ( set -x ; podman  build -t rainbow --build-arg port=8088 . && { podman stop rainbow ; podman rm rainbow ; podman run -p 0.0.0.0:8088:8088 -d --name rainbow rainbow && sleep 1 && podman logs --follow rainbow ; } )


# --------------------------------------------------------------------
# Arguments to control server address (CORS, API, front-end)
ARG addr=http://localhost
ARG port=8884

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
     .env                       ./

COPY frontend/public public
COPY frontend/src    src

ARG addr
ARG port

RUN set -x                                                     &&\
    ls -lA                                                     &&\
    sed -e "s|^VITE_API_ADDR=.*|VITE_API_ADDR=$addr|"            \
        -e "s|^VITE_API_PORT=.*|VITE_API_PORT=$port|" -i .env  &&\
    head .env                                                  &&\
    yarn build                                                 &&\
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
RUN set -x                                                               &&\
    ls -lA                                                               &&\
    export GOOS=linux                                                    &&\
    export CGO_ENABLED=0                                                 &&\
    export GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw"  &&\
    export GOLDFLAGS="-linkmode=external -s -w"                          &&\
    go build ./cmd/server                                                &&\
    ls -sh server                                                        &&\
    ldd server                                                           &&\
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

ARG addr
ARG port

COPY --chown=5505:5505 --from=integrator /target /

# Run as unprivileged
USER teal:teal

# Use UTC time zone by default
ENV TZ        UTC0
ENV WWW_DIR   /var/www
ENV MAIN_ADDR "$addr"
ENV MAIN_PORT "$port"
ENV EXP_PORT  9868

EXPOSE "$port"
EXPOSE 9868

ENTRYPOINT ["/server"]
