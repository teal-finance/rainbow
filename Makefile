.PHONY: help
help:
	# Usage: 'make <target>' where <target> is one of:
	#
	# buildall     Build the backend + frontend
	# build        Build the backend only
	# buildfront   Build the frontend only
	# run          Run the backend
	# runfront     Run the frontend in dev mode
	#
	# container-build  Build the container image (docker or podman or buildash)
	# container-run    Run the container (before: build it if needed)
	# container-stop   Stop the container (automatically called by the other targets)
	# container-rm     Remove the container image (before: stop it if still running)

.PHONY: buildall
buildall: build buildfront

.PHONY: build
build:
	go build ./cmd/server

.PHONY: buildfront
buildfront:
	cd frontend && yarn && yarn build

.PHONY: run
run:
	./server

.PHONY: runfront
runfront:
	cd frontend && yarn dev

##########  Container targets  ##########

expose=1111
addr=http://localhost:$(expose)
port=2222
base=/

.PHONY: container-build
container-build: container-stop
	# Build the container image
	docker  build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow . || \
	podman  build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow . || \
	buildah build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow .

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

.PHONY: container-stop
container-stop:
	# Check if the command "docker" or "podman" is present
	command -v docker || \
	command -v podman

	# Stop the container (if currently running)
	-docker stop   rainbow || \
	podman stop -i rainbow

.PHONY: container-rm
container-rm: container-stop
	# Remove the container image
	docker  rmi rainbow || \
	podman  rmi rainbow
