# Build:
#
#    docker build . -t=rainbow
#    podman build . -t=rainbow
#
# Run with disabled export port:
#
#    docker run -p 0.0.0.0:1234:1234 -d -e API_PORT=1234 -e EXP_PORT=0 --name rainbow rainbow
#    podman run -p 0.0.0.0:1234:1234 -d -e API_PORT=1234 -e EXP_PORT=0 --name rainbow rainbow

FROM docker.io/golang:1.17-alpine3.14 AS builder

WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN go version
RUN go mod download
RUN go mod verify

COPY cmd cmd
COPY pkg pkg
COPY structures.go structures.go

RUN ls -l

# "-s -w" removes all debug symbols https://pkg.go.dev/cmd/link
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -v -mod=readonly ./cmd/rainbow

# smoke test
RUN ls -sh rainbow
RUN ./rainbow -help

WORKDIR /target

RUN cp -a /code/rainbow           .

# HTTPS root certificates (adds about 200 KB)
RUN mkdir -p                                 etc/ssl/certs
RUN cp -a /etc/ssl/certs/ca-certificates.crt etc/ssl/certs

# User & group files
RUN echo 'teal:x:5505:5505::/:' > etc/passwd
RUN echo 'teal:x:5505:'         > etc/group

# --------------------------------------------------------------------
FROM scratch AS final

COPY --chown=5505:5505 --from=builder /target /

# Run as unprivileged
USER teal:teal

ENV TZ   UTC0

ENV API_PORT 1234
ENV EXP_PORT 5678

EXPOSE 1234
EXPOSE 5678

ENTRYPOINT ["/rainbow"]
