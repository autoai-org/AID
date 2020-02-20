package main

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	_ "github.com/mattn/go-sqlite3"
)

func initDatabase() {
	db := storage.GetDefaultDB()
	db.Connect()
	db.CreateTables()
	logger.Info("Successfully initialized database")
}

// put everything (logs, packages, etc) under ~/.autoai/.aid
func initFolder() {
	utilities.InitFolders()
	logger.Info("Successfully created required folders")
}

// insert default values to the db
func addDefaultToDB() {
	entities.SaveInitLogObject()
	logger.Info("Successfully add required objects into database")
}
