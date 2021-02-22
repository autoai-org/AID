// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package workflow

import "github.com/autoai-org/aid/internal/runtime/docker"

// CreateContainer creates a stopped container
// The container can be then started by ```aid start```
func CreateContainer(imageUID string, hostPort string) {
	docker.Create(imageUID, hostPort)
}

// StartContainer starts a stopped container
func StartContainer(containerUID string) {
	docker.Start(containerUID)
}
