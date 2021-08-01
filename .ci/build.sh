#!/bin/bash
echo ">>> Building on CI Platform"
platform=$(uname)
if [[ $platform == "Linux" ]]; then 
  cd components/cmd/tui && go build -ldflags "-s -w -X" main.Version="$(BUILDVERSION)" -X main.Build="$(BUILDTIME)"  -o aid-linux
fi