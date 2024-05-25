.DEFAULT_GOAL := build

# Targets
.PHONY: fmt vet build clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build 

clean:
	go clean
