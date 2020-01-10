// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"path/filepath"
)

var (
	// LogPath tells where the log file should be
	LogPath = filepath.Join(GetBasePath(), "logs")
)
