// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"time"
)

// Log defines the struct of log file
type Log struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Filepath  string    `db:"filepath"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	From      string    `db:"from"`
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
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(l)
}

// NewLogObject creates a new log object and saves to database
func NewLogObject(title string, filepath string, from string) Log {
	log := Log{Title: title, Filepath: filepath, From: from}
	log.Save()
	return log
}
