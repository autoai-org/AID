default:
	@echo "Make sure you are inside components/cmd. This Makefile is ONLY for CMD"
	@echo "Usage:"
	@echo "\tmake test"
	@echo "\tmake build"

test:
	go test -race -coverprofile c.out -v ./...
	
build:
	bash ./build.sh