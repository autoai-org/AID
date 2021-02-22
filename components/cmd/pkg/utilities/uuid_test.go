// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUUIDv4(t *testing.T) {
	uuid := GenerateUUIDv4()
	assert.IsType(t, "string", uuid)
}
