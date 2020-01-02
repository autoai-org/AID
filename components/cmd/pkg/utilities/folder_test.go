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
