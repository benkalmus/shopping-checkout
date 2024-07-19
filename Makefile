.PHONY: vet test race build run clean version help

include vsn.mk
PROJECT_NAME="shopping-checkout"

vet:
	go vet

test: 
	go test -v

race:
	go test -race -v *.go

build:
	@if [ ! -d bin ]; then mkdir bin; fi
	go build -o bin/$(PROJECT_NAME)

run: build
	./bin/$(PROJECT_NAME)

clean: 
	rm -rf bin

version:
	@echo $(PROJECT_VERSION)

help:
	@echo "Usage: make [command]"
	@echo
	@echo "Available commands:"
	@echo "  vet"
	@echo "  test"
	@echo "  race"
	@echo "  build"
	@echo "  run"
	@echo "  clean"
	@echo "  version"