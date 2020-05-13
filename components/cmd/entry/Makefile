BUILDTIME := $(shell date +%FT%T%Z)
default:
	@echo "This makefile is for development purpose only"
	@echo "Usage:"
	@echo "\tmake build  \t Build for local arch"

build:
	go build -ldflags "-s -w -X main.Version="dev" -X main.Build=$(BUILDTIME)"  -o aid
	rm /usr/local/bin/aid
	cp ./aid /usr/local/bin/aid

format:
	gofmt -l -s -w *.go