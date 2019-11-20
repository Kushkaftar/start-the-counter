.PHONY: build
build:
		go build -v ./cmd/sup
		
.DEFAULT_GOAL := build