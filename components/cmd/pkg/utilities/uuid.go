// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	guuid "github.com/google/uuid"
)

// GenerateUUIDv4 returns uuidv4
func GenerateUUIDv4() string {
	id := guuid.New()
	return id.String()
}
