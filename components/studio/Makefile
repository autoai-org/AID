# Copyright (c) 2021 Xiaozhe Yao et al.
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
OSFLAG 				:=
ifeq ($(shell uname -s),Linux)
	OSFLAG +=-linux
endif
ifeq ($(shell uname -s),Darwin)
	OSFLAG +=-macOS
endif

prepare-binaries:
	mkdir src-tauri/binaries
	cp ../cmd/tui/aid$(strip $(OSFLAG)) src-tauri/binaries/aid
	node scripts/append.js

build-binaries:
	rm -r src-tauri/binaries
	make prepare-binaries
	yarn build
	yarn tauri build

build-binaries-debug:
	rm -r src-tauri/binaries
	make prepare-binaries
	yarn build
	yarn tauri build --debug