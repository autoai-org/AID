// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"path/filepath"
	"time"
)

// Log defines the struct of log file
type Log struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Filepath  string    `db:"filepath"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Source    string    `db:"source"`
}

// TableName defines the tablename in database
func (l *Log) TableName() string {
	return "log"
}

// PK defines the primary key of Log
func (l *Log) PK() string {
	return "id"
}

// Save stores log into database
func (l *Log) Save() error {
	l.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(l)
}

// NewLogObject creates a new log object and saves to database
func NewLogObject(title string, filepath string, source string) Log {
	log := Log{Title: title, Filepath: filepath, Source: source}
	return log
}

// FetchLogs returns all logs file
func FetchLogs() []Log {
	logsPointers := make([]*Log, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&logsPointers)
	logs := make([]Log, len(logsPointers))
	for i := range logsPointers {
		logs[i] = *logsPointers[i]
	}
	return logs
}

// DeleteLog removes the log entry from database
func DeleteLog(id string) error {
	log := Log{ID: id}
	db := storage.GetDefaultDB()
	err := db.Delete(&log)
	if err != nil {
		logger.Error("Cannot fetch object")
		logger.Error(err.Error())
	}
	return err
}

// GetLog returns the log object by log id
func GetLog(id string) Log {
	log := Log{ID: id}
	db := storage.GetDefaultDB()
	err := db.FetchOne(&log)
	if err != nil {
		logger.Error("Cannot fetch object")
		logger.Error(err.Error())
	}
	return log
}

// SaveInitLogObject saves the system logs into database
// It should be called when use init to initialize the system
func SaveInitLogObject() {
	prelimaryLogs := [1]string{"system"}
	for _, each := range prelimaryLogs {
		requiredLogPath := filepath.Join(utilities.GetBasePath(), "logs", "system")
		logObj := Log{Title: each, Filepath: requiredLogPath, Source: "Default"}
		logObj.Save()
	}
}
