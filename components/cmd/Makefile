# Copyright (c) 2021 Xiaozhe Yao et al.
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

default:
	@echo -e "==========================================="
	@echo -e "==== This Makefile is for AID TUI Only ===="
	@echo -e "==========================================="
	@echo -e "Usage:"
	@echo -e -e "\tmake format\t\t Format all go code"
	@echo -e "\tmake generate-bins\t Ensemble binary assets"
	@echo -e "\tmake generate-ent\t Generate entities"
	@echo -e "\tmake build-cli\t\t Build binary executable"
	@echo -e "\tmake xbuild\t\t Build binary executable for linux/macOS"

BUILDTIME := $(shell date +%FT%T%Z)
BUILDVERSION := $(shell git rev-parse --short HEAD)
OSFLAG 				:=
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OSFLAG += -linux
endif
ifeq ($(UNAME_S),Darwin)
	OSFLAG += -macOS
endif

format:
	@gofmt -l -s -w **/**/*.go

generate-ent:
	@echo -e ">>> Generating entities..."
	@cd ent && go generate

generate-gql:
	@echo -e ">>> Generating graphql bindings..."
	@go generate

generate-all:
	@make generate-ent
	@make generate-gql

describe-ent:
	@cd ent && go run entgo.io/ent/cmd/ent describe ./schema

build-cli:
	@echo -e ">>> Building CLI utilities..."
	@cd tui && go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o aid
	@cd tui && mv aid ../

build-all-cli:
	@make generate-all
	@make build-cli

ci-build:
	@ echo "Building on "$(OSFLAG)
	@go mod download
	@go mod verify
	@make generate-all 
	cd tui && go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o "aid$(OSFLAG)"


xbuild:
	@echo -e ">>> Starting cross-platform building..."
	@echo -e ">>> This command is supposed to be running on CI Server"
	@echo -e ">>> Make sure you know what you are doing:)"
	@go mod download
	@go mod verify
	@make generate-all
	@cd tui && env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o aid-macOS
	@cd tui && env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o aid-linux
	@cd tui && env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.Version="$(BUILDVERSION)" -X main.Build=$(BUILDTIME)"  -o aid-windows