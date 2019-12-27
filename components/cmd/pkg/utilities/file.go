// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"io/ioutil"
)

var logger = NewLogger()

// ReadFileContent returns the content of filename
func ReadFileContent(filename string) string {
	var content string
	byteContent, err := ioutil.ReadFile(filename)
	if err != nil {
		content = "Read " + filename + " Failed!"
	} else {
		content = string(byteContent)
	}
	return content
}

// WriteContentToFile writes contents to file
func WriteContentToFile(filepath string, content string) error {
	contentInByte := []byte(content)
	err := ioutil.WriteFile(filepath, contentInByte, 0644)
	if err != nil {
		logger.Error("Cannot write to file " + filepath)
		return err
	}
	return nil
}
