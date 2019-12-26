package daemon

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

var logger = utilities.NewLogger()

// RunServer starts the http service
func RunServer(port string) {
	logger.Info("Initiating Service...")
	r := getRouter()
	r.Run("0.0.0.0:" + port)
}
