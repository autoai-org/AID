package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

func initDatabase() {
	db := storage.GetDefaultDB()
	db.Connect()
	db.CreateTables()
}

// put everything (logs, packages, etc) under ~/.autoai/.aid
func initFolder() {
	utilities.InitFolders()
}