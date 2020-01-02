// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWriteContent(t *testing.T) {
	testString := "test"
	err := WriteContentToFile("./test.test", testString)
	assert.Nil(t, err)
	os.Remove("./test.test")
}

func TestReadFileContent(t *testing.T) {
	testString := "test"
	err := WriteContentToFile("./test.test", testString)
	assert.Nil(t, err)
	content := ReadFileContent("./test.test")
	assert.Equal(t, content, testString)
	os.Remove("./test.test")
}
