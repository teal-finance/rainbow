# Build:
#
#    docker build . -t=rainbow
#    podman build . -t=rainbow
#
# Run with disabled export port:
#
#    docker run -p 0.0.0.0:1234:1234 -d -e API_PORT=1234 -e EXP_PORT=0 --name rainbow rainbow
#    podman run -p 0.0.0.0:1234:1234 -d -e API_PORT=1234 -e EXP_PORT=0 --name rainbow rainbow





# ---- Build -----
# DOCKER_BUILDKIT=1
# docker build --file build/Dockerfile.local . --tag tealapi

# ---- Run ----
# export IDB_URL=https://db.n20.hessling.fr:41444
# export IDB_ORG=j3
# export IDB_TOKEN=lwEperkgPNLZQPmidUQFjNxOrgWJ4UcG47lKPSrw9W7Kls9t1YW4rpZhisBewbxF8EyjGEIThoLgfvegdmlA9A==
# docker run -e IDB_URL -e IDB_ORG -e IDB_TOKEN              \
#            -v "/etc/teal:/etc/teal:Z"                      \
#            -v "/var/cache/teal:/var/cache/teal:delegate,Z" \
#            -v "/var/www/teal.finance:/var/www:Z"           \
#            -rm --name tealapi tealapi

# --------------------------------------------------------------------
FROM docker.io/node:16-alpine3.14 AS web_builder
RUN node    --version
RUN yarnpkg --version

WORKDIR /code

COPY package.json .
COPY yarn.lock    .

RUN yarnpkg install --frozen-lockfile

COPY .eslintrc.json     .
COPY index.html         .
COPY postcss.config.js  .
COPY tailwind.config.js .
COPY tsconfig.json      .
COPY vite.config.ts     .

COPY public public
COPY src    src

RUN ls -l

RUN yarnpkg build
RUN yarnpkg compress

# --------------------------------------------------------------------
FROM docker.io/golang:1.17-alpine3.14 AS go_builder

WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN go version
RUN go mod download
RUN go mod verify

COPY cmd cmd
COPY pkg pkg
COPY structures.go .

RUN ls -l

# "-s -w" removes all debug symbols https://pkg.go.dev/cmd/link
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -v -mod=readonly ./cmd/rainbow

# smoke test
RUN ls -sh rainbow
RUN ./rainbow -help

# --------------------------------------------------------------------
FROM docker.io/golang:1.17-alpine3.14 AS integrator

WORKDIR /target

# HTTPS root certificates (adds about 200 KB)
RUN mkdir -p                                 etc/ssl/certs
RUN cp -a /etc/ssl/certs/ca-certificates.crt etc/ssl/certs

# User & group files
RUN echo 'teal:x:5505:5505::/:' > etc/passwd
RUN echo 'teal:x:5505:'         > etc/group

# Back-end and static website
COPY --from=go_builder  /code/rainbow .
COPY --from=web_builder /code/dist    var/www

# --------------------------------------------------------------------
FROM scratch AS final

COPY --chown=5505:5505 --from=integrator /target /

# Run as unprivileged
USER teal:teal

ENV TZ   UTC0

ENV API_PORT 1234
ENV EXP_PORT 5678
ENV WWW_DIR  /var/www

EXPOSE 1234
EXPOSE 5678

ENTRYPOINT ["/rainbow"]
