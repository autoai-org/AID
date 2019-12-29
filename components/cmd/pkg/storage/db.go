// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package storage

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/ilibs/gosql/v2"
)

var logger = utilities.NewDefaultLogger("./logs/system.log")

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
	config := utilities.GetDefaultConfig()
	db := NewDB(config.Read("db_driver"), config.Read("db_uri"))
	db.Connect()
	return db
}

// CreateTables Create all Tables
func (db *Database) CreateTables() {
	// Read SQL
	var sqlString = utilities.GetRemoteFile("https://raw.githubusercontent.com/autoai-org/CVPM/master/components/cmd/pkg/storage/db.sql")
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
