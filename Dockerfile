FROM docker.io/golang:1.17-buster AS builder

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

RUN CGO_ENABLED=0 go build -v -mod=readonly ./cmd/rainbow

# smoke test
RUN ls -sh rainbow
RUN ./rainbow -help

WORKDIR /target

# Time zones and HTTPS root certificates
RUN mkdir -p   etc/ssl/certs   usr/share
RUN cp -a /usr/share/zoneinfo  usr/share
RUN cp -a /etc/ssl/certs/ca-certificates.crt etc/ssl/certs
RUN cp -a /code/rainbow           .

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
