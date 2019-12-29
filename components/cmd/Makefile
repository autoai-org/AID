default:
	@echo "Make sure you are inside components/cmd. This Makefile is ONLY for CMD"
	@echo "Usage:"
	@echo "\tmake format"
	@echo "\tmake test"
	@echo "\tmake build"

format:
	gofmt -l -s -w **/**/*.go

test:
	go test -race -coverprofile c.out -v ./...
	rm -rf **/**/logs/
	
build:
	bash ./build.sh