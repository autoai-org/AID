// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	assert.Equal(t, os.Stdout, logger.Out)
}
