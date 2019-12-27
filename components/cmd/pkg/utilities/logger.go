// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"github.com/sirupsen/logrus"
	"os"
)

// NewLogger returns the Logger Object
func NewLogger() *logrus.Logger {
	var logger = logrus.New()
	logger.Out = os.Stdout
	return logger
}
