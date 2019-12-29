package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
)

func initDatabase() {
	db := storage.GetDefaultDB()
	db.Connect()
	db.CreateTables()
}