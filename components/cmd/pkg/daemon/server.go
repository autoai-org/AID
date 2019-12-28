// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

var logger = utilities.NewDefaultLogger("./logs/system.log")

// RunServer starts the http service
func RunServer(port string, sslcert string, sslkey string) {
	logger.Info("Initiating Service...")
	r := getRouter()
	err := r.RunTLS("127.0.0.1:"+port, sslcert, sslkey)
	utilities.CheckError(err, "Cannot Start Server")
}
