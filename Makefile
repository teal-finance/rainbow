.PHONY: help
help:
	# Usage: 'make <target>' where <target> is one of:
	#
	# make all -j         Build both backend and frontend
	# make server         Build the backend only
	# make front          Build the frontend only
	# make clean          Clean all
	#
	# make run            Run the backend
	# make run-front      Run the frontend in dev mode
	#
	# make container-run  Build and run the container
	# make container-rm   Stop and remove the container image
	#
	# Example:
	#
	# make clean all -j   Rebuild all in parallel

# Allow using a different Go executable by running as
# "GOEXE=xxx make ..." or as "make ... GOEXE=xxx ..."
GOEXE ?= go

.PHONY: clean
clean:
	rm -fr server frontend/dist

.PHONY: all
all: server frontend/dist

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
	${GOEXE} run ./cmd/server

.PHONY: run-front
run-front:
	cd frontend && yarn dev

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
