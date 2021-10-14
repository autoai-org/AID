// Copyright (c) 2021 Xiaozhe Yao
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"fmt"
	"net"
)

func GetAvailablePort() int {
	// by default, we start at 8080
	port := 8080
	for {
		ln, err := net.Listen("tcp", ":"+fmt.Sprint(port))
		if err == nil {
			break
		} else {
			port = port + 1
		}
		ln.Close()
	}
	return port
}
