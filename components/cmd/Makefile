# Copyright (c) 2021 Xiaozhe Yao et al.
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
default:
	@echo "==========================================="
	@echo "==== This Makefile is for AID TUI Only ===="
	@echo "==========================================="
	@echo "Usage:"
	@echo "\tmake format\t\t Format all go code"
	@echo "\tmake generate-bins\t Ensemble binary assets"
	@echo "\tmake generate-ent\t Generate entities"
	@echo "\tmake build-cli\t\t Build binary executable"

BUILDTIME := $(shell date +%FT%T%Z)
format:
	gofmt -l -s -w **/**/*.go

generate-bins:
	go-bindata -pkg docker -o internal/runtime/docker/blob.go internal/assets/

generate-ent:
	rm -rf ./ent/generated
	mkdir ./ent/generated
	entc generate ./ent/schema --target ./ent/generated 

build-cli:
	# make generate-bins
	cd tui && go build -ldflags "-s -w -X main.Version="dev" -X main.Build=$(BUILDTIME)"  -o aid
	cd tui && mv aid ../