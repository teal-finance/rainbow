# 🏎️ Run Rainbow from Golang source code

## 📑 Requirements

- Go v1.21+

On Debian/Ubuntu, the command `sudo apt install golang` may install an older version.
Check it with `apt list --all-versions golang golang-1.*`.

To install a more recent Go version, you may try:

    sudo apt purge   golang
    sudo apt install golang-1.22

[Snap](<https://en.wikipedia.org/wiki/Snap_(package_manager)>) provides a simple way to install the latest Go version on most Linux distributions:

    snap install go   --classic
    go version

## 🪃 Run the CLI

The Rainbow project provides a Command Line Interface (CLI).
The command `./cli` retrieves and prints a pretty nice table of all options.
⚠️ This usually takes more than 10 minutes! Be patient…

    go run ./cmd/cli

![CLI screenshot](../doc/cli.jpg)

## 🛰️ Run the API server

The API server also serves the web frontend.

    go run ./cmd/server

If you have generated the weé static files, you can open <http://localhost:8090>.

This Rainbow server in live on <https://teal.finance/rainbow/>

## 📥 Run the API client

The API client connects to the API server using a valid token (JWT).

    go run ./cmd/client
