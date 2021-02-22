// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

// DefaultLogger is the instance shared by all modules
var DefaultLogger *logrus.Logger

// Verbose sets the logger level.
var Verbose bool

// NewDefaultLogger returns the Logger Object
func NewDefaultLogger() *logrus.Logger {
	if DefaultLogger != nil {
		return DefaultLogger
	}
	logPath := filepath.Join(GetBasePath(), "logs", "system")
	DefaultLogger = NewLogger(logPath)
	return DefaultLogger
}

// NewLogger returns a new Logger
func NewLogger(logPath string) *logrus.Logger {
	var logger *logrus.Logger
	CreateFolderIfNotExist(filepath.Dir(logPath))
	logger = logrus.New()
	if Verbose {
		rotateFileHook, err := NewRotateFileHook(RotateFileConfig{
			Filename:   logPath,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     28, //days
			Level:      logrus.InfoLevel,
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: time.RFC822,
			},
		})
		logger.SetOutput(os.Stdout)
		logger.AddHook(rotateFileHook)
		ReportError(err, "Cannot initialize logger")
	} else {
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		ReportError(err, "Cannot open file")
		logger.SetOutput(file)
	}

	return logger
}

var logger = NewDefaultLogger()
