// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"strings"

	guuid "github.com/google/uuid"
)

// GenerateUUIDv4 returns uuidv4
func GenerateUUIDv4() string {
	id := guuid.New().String()
	firstBlock := strings.Split(id, "-")
	return firstBlock[0]
}
