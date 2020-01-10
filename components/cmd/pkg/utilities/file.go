// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var logger = NewDefaultLogger()

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

// GetRemoteFile returns a string that is included in a remote file
func GetRemoteFile(url string) string {
	resp, err := http.Get(url)
	CheckError(err, "Cannot fetch remote file..., Please check your internet connection!")
	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err, "Cannot read from response..., Please check your internet connection!")
	return string(body)
}

// ReadFileIfModified will return filecontent in byte mode if file has been modified
func ReadFileIfModified(filename string, lastMod time.Time) ([]byte, time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), err
	}
	fileindicator := []byte(filename + "\n")
	msg := append(fileindicator, p...)
	return msg, fi.ModTime(), nil
}
