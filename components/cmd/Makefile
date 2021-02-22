default:
	@echo "Make sure you are inside components/cmd. This Makefile is ONLY for CMD"
	@echo "Usage:"
	@echo "\tmake format\t Format all go code"
	@echo "\tmake test\t Run Go Test"
	@echo "\tmake build  \t Build for local arch"
	@echo "\tmake xbuild \t Build for different platforms"

format:
	gofmt -l -s -w **/**/*.go

test:
	go test -race -coverprofile c.out -v ./...
	rm -rf **/**/logs/
	
build:
	bash ./build.sh

xbuild:
	bash ./xbuild.sh

upload-edge:
	bash ./scripts/upload-edge.sh