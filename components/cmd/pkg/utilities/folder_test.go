package utilities

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIfPathExists(t *testing.T) {
	assert.True(t, IsExists("folder.go"))
	assert.False(t, IsExists("folder_not_exist.go"))
	assert.True(t, IsExists("./"))
	assert.False(t, IsExists("./non_exists"))
}

func TestGetHomeDir(t *testing.T) {
	assert.IsType(t, "string", GetHomeDir())
}

func TestCreateFolderIfNotExists(t *testing.T) {
	assert.False(t, IsExists("./non_exists"))
	CreateFolderIfNotExist("./non_exists")
	assert.True(t, IsExists("./non_exists"))
	os.Remove("./non_exists")
}