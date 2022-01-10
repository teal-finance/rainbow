help:
	# Please use 'make <target>' where <target> is one of:
	#
	#   build          -- build the backend
	#   buildfront     -- build the frontend
	#   buildall       -- build the backend and the frontend
	#   buildimage     -- build the container image (docker or podman on buildash)
	#   rmi            -- remove the container image "rainbow"
	#
	#   run            -- run the backend
	#   runfront       -- run the frontend in dev mode
	#   runcontainer   -- run the container image (docker or podman)
	#   stopcontainer  -- stop the container image

build:
	go build ./cmd/server
.PHONY: build

buildfront:
	cd frontend && yarn && yarn build
.PHONY: buildfront

buildall: build buildfront
.PHONY: buildall

expose=1111
addr=http://localhost:$(expose)
port=2222
base=/

buildimage:
	# Build the container image
	docker  build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow . || \
	podman  build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow . || \
	buildah build --build-arg addr=$(addr) --build-arg port=$(port) --build-arg base=$(base) -t rainbow .
.PHONY: buildimage

rmi: stopcontainer
	# Remove the container image
	docker  rmi rainbow || \
	podman  rmi rainbow
.PHONY: rmi

run:
	./server
.PHONY: run

runfront:
	cd frontend && yarn dev
.PHONY: run

runcontainer: stopcontainer buildimage
	# Start the container
	docker run -d --rm -p 0.0.0.0:$(expose):$(port) -e EXP_PORT=9868 --name rainbow rainbow || \
	podman run -d --rm -p 0.0.0.0:$(expose):$(port) -e EXP_PORT=9868 --name rainbow rainbow

	# Open the web-browser
	xdg-open http://localhost:$(expose)$(base)

	# Print backend logs
	docker logs --follow rainbow || \
	podman logs --follow rainbow
.PHONY: runcontainer

stopcontainer:
	# Stop the container (if currently running)
	docker stop -i rainbow || \
	podman stop -i rainbow
.PHONY: runcontainer
