package utilities

import (
	"os"
	"github.com/stretchr/testify/assert"
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