.PHONY: help
help:
	# Usage: 'make <target>' where <target> is one of:
	#
	#   all            Build both backend and frontend
	#   server         Build the backend only
	#   front          Build the frontend only
	#   clean          Clean all
	#
	#   run            Run the backend
	#   run-front      Run the frontend in dev mode
	#
	#   container-run  Build and run the container
	#   container-rm   Stop and remove the container image
	#
	# Example:
	#
	#   make clean build -j2       Rebuild all in parallel

.PHONY: clean
clean:
	rm -fr server frontend/dist

.PHONY: all
all: server frontend/dist

server:
	go build ./cmd/server

.PHONY: front
front: frontend/dist

frontend/dist: frontend/src/*
	cd frontend && yarn && yarn build && yarn compress

.PHONY: run
run:
	go run ./cmd/server

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
