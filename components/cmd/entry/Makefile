BUILDTIME := $(shell date +%FT%T%Z)
default:
	@echo "This makefile is for development purpose only"
	@echo "Usage:"
	@echo "\tmake build  \t Build for local arch"

build:
	go build -ldflags "-s -w -X main.Version="Dev" -X main.Build=$(BUILDTIME)"  -o aid
