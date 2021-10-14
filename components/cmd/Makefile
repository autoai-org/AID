# Copyright (c) 2021 Xiaozhe Yao et al.
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

default:
	@echo  "==========================================="
	@echo  "==== This Makefile is for AID TUI Only ===="
	@echo  "==========================================="
	@echo  "Usage:"
	@echo   "\tmake format\t\t Format all go code"
	@echo  "\tmake generate-bins\t Ensemble binary assets"
	@echo  "\tmake generatent\t Generate entities"
	@echo  "\tmake build-cli\t\t Build binary executable"
	@echo  "\tmake xbuild\t\t Build binary executable for linux/macOS"

BUILDTIME := $(shell date +%FT%T%Z)
BUILDVERSION := $(shell git rev-parse --short HEAD)
OSFLAG 				:=
ifeq ($(shell uname -s),Linux)
	OSFLAG +=-linux
endif
ifeq ($(shell uname -s),Darwin)
	OSFLAG +=-macOS
endif


format:
	@gofmt -l -s -w **/**/*.go

generatent:
	@echo  ">>> Generating entities..."
	@cd ent && go generate

generate-gql:
	@echo  ">>> Generating graphql bindings..."
	@go generate

build-studio-site:
	cd ../studio && yarn && yarn build

cp-studio-site:
	cp -r ../studio/build internal/daemon/

generate-all:
	@make generatent
	@make generate-gql

describent:
	@cd ent && go run entgo.io/ent/cmd/ent describe ./schema

build-cli:
	@go env -w CC=gcc CXX="g++"
	@echo  ">>> Building CLI utilities..."
	@cd tui && go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o aid
	@cd tui && mv aid ../

build-all-cli:
	@go env -w CC=gcc CXX="g++"
	@make generate-all
	@make build-cli

ci-build:
	@echo "Building on "$(OSFLAG)
	@go mod download
	@go mod verify
	@make generate-all 
	@make build-studio-site
	@make cp-studio-site
	cd tui && go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o "aid$(strip $(OSFLAG))"

xbuild:
	@echo  ">>> Starting cross-platform building..."
	@echo  ">>> This command is supposed to be running on CI Server with cross-platform toolkits"
	@echo  ">>> Make sure you know what you are doing:)"
	@go mod download
	@go mod verify
	@make generate-all
	@cd tui && env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o aid-macOS
	@cd tui && env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o aid-linux
	@cd tui && env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o aid-windows

install:
	cp ./aid /usr/local/bin
