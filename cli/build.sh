#/bin/bash
env GOOS=windows GOARCH=amd64 go build . -o dist/cvpm-windows
env GOOS=linux GOARCH=amd64 go build . -o dist/cvpm-linux
env GOOS=darwin GOARCH=amd64 go build 