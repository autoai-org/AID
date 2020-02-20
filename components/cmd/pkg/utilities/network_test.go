// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOutboundIP(t *testing.T) {
	ip := GetOutboundIP().String()
	assert.IsType(t, "string", ip)
}
