# Build:
#
#    docker  build -t rainbow .
#    podman  build -t rainbow .
#    buildah build -t rainbow .
#
# Run:
#
#    docker run -p 0.0.0.0:8090:8090 -d -e MAIN_PORT=8090 -e EXP_PORT=9868 --name rainbow rainbow
#    podman run -p 0.0.0.0:8090:8090 -d -e MAIN_PORT=8090 -e EXP_PORT=9868 --name rainbow rainbow
#
# One single command line:
#
#    ( set -x ; podman  build -t rainbow . && { podman stop rainbow && podman rm rainbow ; podman run -p 0.0.0.0:8090:8090 -d -e MAIN_PORT=8090 -e EXP_PORT=9868 --name rainbow rainbow && sleep 1 && podman logs --follow rainbow ; } )

# --------------------------------------------------------------------
FROM docker.io/node:16-alpine3.14 AS web_builder

WORKDIR /code

COPY frontend/package.json \
     frontend/yarn.lock   ./

RUN set -x                          &&\
    node --version                  &&\
    yarn --version                  &&\
    yarn install --frozen-lockfile

COPY frontend/.eslintrc.json     \
     frontend/index.html         \
     frontend/postcss.config.js  \
     frontend/tailwind.config.js \
     frontend/tsconfig.json      \
     frontend/vite.config.ts     \
     .env                       ./
COPY frontend/public public
COPY frontend/src    src

RUN set -x                                        &&\
    ls -lA                                        &&\
    find . -name '*.env' -print -exec head {} +   &&\
    yarn build                                    &&\
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
RUN set -x                                                                &&\
    ls -lA                                                                &&\
    export GOOS=linux                                                     &&\
    export CGO_ENABLED=1                                                  &&\
    export GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw"   &&\
    export GOLDFLAGS="-linkmode=external -s -w"                           &&\
    CGO_ENABLED=0 go build -mod=readonly ./cmd/server                     &&\
    ls -sh server                                                         &&\
    ldd server                                                            &&\
    ./server -help  # smoke test

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

COPY --chown=5505:5505 --from=integrator /target /

# Run as unprivileged
USER teal:teal

# Use UTC time zone by default
ENV TZ        UTC0
ENV WWW_DIR   /var/www
ENV MAIN_PORT 8090
ENV EXP_PORT  9868

EXPOSE 8090
EXPOSE 9868

ENTRYPOINT ["/server"]
