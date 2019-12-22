package utilities

import (
	"github.com/sirupsen/logrus"
	"os"
)

// NewLogger returns the Logger Object
func NewLogger () *logrus.Logger {
	var logger = logrus.New()
	logger.Out = os.Stdout
	return logger
}
