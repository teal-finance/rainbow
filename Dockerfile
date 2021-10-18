# Build:
#
#    docker build . -t=rainbow
#    podman build . -t=rainbow
#
# Run with disabled export port:
#
#    docker run -p 0.0.0.0:1234:1234 -d -e MAIN_PORT=1234 -e EXP_PORT=0 --name rainbow rainbow
#    podman run -p 0.0.0.0:1234:1234 -d -e MAIN_PORT=1234 -e EXP_PORT=0 --name rainbow rainbow

# --------------------------------------------------------------------
FROM docker.io/node:16-alpine3.14 AS web_builder
RUN node    --version
RUN yarnpkg --version

WORKDIR /code

COPY frontend/package.json .
COPY frontend/yarn.lock    .

RUN yarnpkg install --frozen-lockfile

COPY frontend/.eslintrc.json     .
COPY frontend/index.html         .
COPY frontend/postcss.config.js  .
COPY frontend/tailwind.config.js .
COPY frontend/tsconfig.json      .
COPY frontend/vite.config.ts     .

COPY frontend/public frontend/public
COPY frontend/src    frontend/src

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

ENV MAIN_PORT 1234
ENV EXP_PORT 5678
ENV WWW_DIR  /var/www

EXPOSE 1234
EXPOSE 5678

ENTRYPOINT ["/rainbow"]
