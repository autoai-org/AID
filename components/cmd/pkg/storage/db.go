// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package storage

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/ilibs/gosql/v2"

	// import sqlite3 as driver for database
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var logger = utilities.NewDefaultLogger()

// Database is the primary class for persisting contents
type Database struct {
	configs     map[string]*gosql.Config
	isConnected bool
}

// DefaultDB is the shared Database Instance shared by all modules
var DefaultDB *Database

// NewDB returns a new Database Instance
func NewDB(driver string, uri string) *Database {
	configs := make(map[string]*gosql.Config)
	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  driver,
		Dsn:     uri,
		ShowSql: true,
	}
	return &Database{
		configs:     configs,
		isConnected: false,
	}
}

// GetDefaultDB returns the default DB instance
func GetDefaultDB() *Database {
	if DefaultDB != nil {
		DefaultDB.Connect()
		return DefaultDB
	}
	// TODO: the config should be from config.json
	//db := NewDB(config.Read("db_driver"), config.Read("db_uri"))
	dbURI := filepath.Join(utilities.GetBasePath(), "aid.db")
	db := NewDB("sqlite3", dbURI)
	db.Connect()
	return db
}

// CreateTables Create all Tables
func (db *Database) CreateTables() {
	// Read SQL
	//var sqlString = utilities.GetRemoteFile("https://raw.githubusercontent.com/autoai-org/CVPM/master/components/cmd/pkg/storage/db.sql")
	var sqlString = `
	--- Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
	---
	--- This software is released under the MIT License.
	--- https://opensource.org/licenses/MIT
	/* Create required tables */
	CREATE TABLE IF NOT EXISTS package (
		id TEXT PRIMARY KEY,
		name TEXT,
		localpath TEXT,
		vendor TEXT,
		status TEXT,
		created_at DATETIME,
		updated_at DATETIME,
		remote_url TEXT
	);
	
	CREATE TABLE IF NOT EXISTS event (
		id TEXT PRIMARY KEY,
		title TEXT,
		data TEXT,
		source TEXT,
		status TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS log (
		id TEXT PRIMARY KEY,
		title TEXT,
		filepath TEXT,
		source TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS solver (
		id TEXT PRIMARY KEY,
		package TEXT,
		vendor TEXT,
		name TEXT,
		solverpath TEXT,
		imagename TEXT,
		created_at DATETIME,
		updated_at DATETIME,
		status TEXT
	);
	
	CREATE TABLE IF NOT EXISTS webhook (
		id TEXT PRIMARY KEY,
		remote_url TEXT,
		created_at DATETIME,
		updated_at DATETIME,
		status TEXT,
		event TEXT
	);
	
	CREATE TABLE IF NOT EXISTS apikey (
		id TEXT PRIMARY KEY,
		aidkey TEXT,
		created_at DATETIME,
		updated_at DATETIME,
		status TEXT
	);
	
	CREATE TABLE IF NOT EXISTS privateenvironment (
		id TEXT PRIMARY KEY,
		name TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS environmentvariable (
		id TEXT PRIMARY KEY,
		envkey TEXT,
		envvalue TEXT,
		environment TEXT,
		PackageId TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS image (
		id TEXT PRIMARY KEY,
		name TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS container (
		id TEXT PRIMARY KEY,
		name TEXT,
		imageid TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS runningsolver (
		id TEXT PRIMARY KEY,
		imageid TEXT,
		status TEXT,
		imagename TEXT,
		entrypoint TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	
	CREATE TABLE IF NOT EXISTS dataset (
		id TEXT PRIMARY KEY,
		name TEXT,
		localpath TEXT,
		status TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	`
	gosql.Exec(sqlString)
}

// Connect tries to connect the database
func (db *Database) Connect() (err error) {
	gosql.SetDefaultLink("default")
	err = gosql.Connect(db.configs)
	if err == nil {
		db.isConnected = true
	}
	return err
}

// Insert saves object into database
func (db *Database) Insert(obj interface{}) (err error) {
	_, err = gosql.Model(obj).Create()
	return err
}

// Fetch returns all objects in the database
func (db *Database) Fetch(obj interface{}) (err error) {
	err = gosql.Model(obj).All()
	return err
}

// FetchOne returns the only object by requested obj
// It should be used only when the requested column is unique
func (db *Database) FetchOne(obj interface{}) (err error) {
	err = gosql.Model(obj).Get()
	return err
}

// Delete removes object from database
func (db *Database) Delete(obj interface{}) (err error) {
	_, err = gosql.Model(obj).Delete()
	return err
}

// Update updates the object in the database
func (db *Database) Update(obj interface{}) (err error) {
	_, err = gosql.Model(obj).Update()
	return err
}

// Count returns the count of the requested object
func (db *Database) Count(obj interface{}) (count int64, err error) {
	count, err = gosql.Model(obj).Count()
	return count, err
}
