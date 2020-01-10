// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

// DefaultLogger is the instance shared by all modules
var DefaultLogger *logrus.Logger

// NewDefaultLogger returns the Logger Object
func NewDefaultLogger() *logrus.Logger {
	if DefaultLogger != nil {
		return DefaultLogger
	}
	logPath := filepath.Join(GetBasePath(), "logs", "system")
	return NewLogger(logPath)
}

// NewLogger returns a new Logger
func NewLogger(logPath string) *logrus.Logger {
	var logger *logrus.Logger
	CreateFolderIfNotExist(filepath.Dir(logPath))
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
	if err != nil {
		logrus.Error("Cannot Initialize Logger..")
		logrus.Error(err.Error())
	}
	logger = logrus.New()
	logger.AddHook(rotateFileHook)
	return logger
}
