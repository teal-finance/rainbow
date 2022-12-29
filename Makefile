.PHONY: help
help:
	# make all -j         Build both backend and frontend
	# make server         Build the backend only
	# make front          Build the frontend only
	# make clean          Clean all
	#
	# make run            Run the backend in dev mode
	# make run-ui         Run the frontend in dev mode
	#
	# make test           (Go only) Test the backend
	# make cov            (Go only) Browse test coverage
	# make fmt            (Go only) Generate code, Format code, Check build
	# make vet            (Go only) Generate, Format, Build, Test, Lint and Run commands
	#
	# make container-run  Build and run the container
	# make container-rm   Stop and remove the container image
	#
	# Examples:
	#
	# make clean all -j   Rebuild all in parallel
	# make fmt test vet   Update/verify backend before commit

# Allow using a different Go executable by running as
# "GOEXE=xxx make ..." or as "make ... GOEXE=xxx ..."
GOEXE ?= go

.PHONY: clean
clean:
	rm -fr frontend/dist server code-coverage.out

.PHONY: all
all: frontend/dist server

server: go.sum
	${GOEXE} build -o $@ ./cmd/server

go.mod:
	${GOEXE} mod tidy
	${GOEXE} mod verify

go.sum: go.mod

.PHONY: front
front: frontend/dist

frontend/dist: frontend/src/*
	cd frontend && \
	{ yarn    && yarn    build && yarn    compress; } || \
	{ yarnpkg && yarnpkg build && yarnpkg compress; }

.PHONY: run
run: go.sum
	${GOEXE} run -race ./cmd/server -dev

.PHONY: run-ui
run-ui:
	cd frontend && yarn dev

##########  Backend only  ##########

.PHONY: gen
fmt:
	go mod tidy
	go generate ./...
	go run mvdan.cc/gofumpt@latest -w -extra -l -lang 1.19 .
	go build ./...

.PHONY: test
test:
	go test -race -tags=rainbow -coverprofile=code-coverage.out ./...

.PHONY:
cov: test
	go tool cover -html code-coverage.out

.PHONY: vet
vet:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix || true
	go run -race ./cmd/client || true
	go run -race ./cmd/cli
	timeout 16m go run -race ./cmd/server  || echo "Terminated by timeout 16m"

##########  Container targets  ##########

# Below default values can be changes using:
# make container-run expose=3333 port=3333
expose=1111
addr=http://localhost:$(expose)
port=2222
base=/

# Build the container image (docker or podman or buildash)
.PHONY: container-build
container-build: container-stop
	# Build the container image
	docker  build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow . || \
	podman  build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow . || \
	buildah build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow .

# Build and run the container
.PHONY: container-run
container-run: container-build
	# Start the container
	docker run -d --rm -p 0.0.0.0:$(expose):$(port) -e EXP_PORT=9868 --name rainbow rainbow || \
	podman run -d --rm -p 0.0.0.0:$(expose):$(port) -e EXP_PORT=9868 --name rainbow rainbow

	# Open the web-browser
	xdg-open http://localhost:$(expose)$(base)

	# Print backend logs
	docker logs --follow rainbow || \
	podman logs --follow rainbow

# Stop the container (automatically called by the other targets)
.PHONY: container-stop
container-stop:
	# Check if the command "docker" or "podman" is present
	command -v docker || \
	command -v podman

	# Stop the container (if currently running)
	-docker stop   rainbow || \
	podman stop -i rainbow

# Stop the container if still running, then remove its image
.PHONY: container-rm
container-rm: container-stop
	# Remove the container image
	docker  rmi rainbow || \
	podman  rmi rainbow
