help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo
	@echo "  build               -- build the backend"
	@echo "  buildfront          -- build the frontend"
	@echo "  buildall            -- build the backend and the frontend"
	@echo ""
	@echo "  run                 -- run the backend"
	@echo "  runfront            -- run the frontend in dev mode"
	@echo

build:
	go build ./cmd/server
.PHONY: build

buildfront:
	cd frontend && yarn build
.PHONY: buildfront

buildall: build buildfront
.PHONY: buildall

run:
	./server
.PHONY: run

runfront:
	cd frontend && yarn dev
.PHONY: run

