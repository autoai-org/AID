# Copyright (c) 2021 Xiaozhe Yao et al.
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

prepare:
	cp ../cmd/aid ./aid-x86_64-apple-darwin
	

build-osx:
	yarn run build
	yarn run build
	cp build/index.html build/200.html
	yarn tauri build
