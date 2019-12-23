package storage

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var logger = utilities.NewLogger()

// Database is the primary class for persisting contents
type Database struct {
	db *sqlx.DB
	dbPath string
}

// New returns a Database instance
func New (schema string, dbPath string) *Database {
	exists := utilities.IsExists(dbPath)
	if !exists {
		logger.WithFields(logrus.Fields{
			"module": "storage",
		  }).Error("File Not Exists " + dbPath)
	}
	database, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "storage",
		  }).Error("Cannot Open Database " + dbPath)
	}
	err = database.Ping()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "storage",
		  }).Error("Cannot Connect to Database " + dbPath)
	}
	return &Database{db: database, dbPath: dbPath}
}

// connect should be run before any queries.
func (db *Database) connect () {

}
