# Copyright (c) 2021 Xiaozhe Yao et al.
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

prepare-binaries:
	mkdir src-tauri/binaries
	mv ../cmd/aid src-tauri/binaries/aid
	node scripts/append.js

build_binaries:
	make prepare-binaries
	yarn tauri build