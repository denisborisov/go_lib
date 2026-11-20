# export GOGC=200
# export GODEBUG=gctrace=1

.DEFAULT_GOAL := build
.PHONY: fmt vet build clean

get:
	go get ./...

mod:
	go mod tidy

fmt: mod
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

test: build
	go test

clean:
	rm ./go_mod
