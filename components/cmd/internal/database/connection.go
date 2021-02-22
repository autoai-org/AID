package database

import (
	"context"
	"path/filepath"

	ent "github.com/autoai-org/aid/ent/generated"
	"github.com/autoai-org/aid/internal/utilities"

	// import sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// DefaultDB is the instance shared by all modules
var DefaultDB *ent.Client

// NewDefaultDB returns the database client
func NewDefaultDB() *ent.Client {
	if DefaultDB != nil {
		return DefaultDB
	}
	client, err := ent.Open("sqlite3", filepath.Join(utilities.GetBasePath(), "aid.db?_fk=1"))
	DefaultDB = client
	utilities.ReportError(err, "cannot open database")
	if err := client.Schema.Create(context.Background()); err != nil {
		utilities.ReportError(err, "Failed updating database schema")
	}
	return DefaultDB
}
