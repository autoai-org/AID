// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/codeclysm/extract"
)

// Uncompress receives a compressed file and a path,
// returns the uncompressed folder path
func Uncompress(compressedFile string, folderPath string) error {
	data, _ := ioutil.ReadFile(compressedFile)
	buffer := bytes.NewBuffer(data)
	return extract.Archive(context.Background(), buffer, folderPath, nil)
}
