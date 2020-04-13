// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/autoai-org/aid/components/cmd/pkg/storage"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
)

// PrivateEnvironment allows users to add new environments other than dev, test, prod
// This feature will not be enabled in a near future
type PrivateEnvironment struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

// EnvironmentVariable stores a list of environment variables inside a environment
type EnvironmentVariable struct {
	ID          string `db:"id"`
	Key         string `db:"envkey"`
	Value       string `db:"envvalue"`
	Environment string `db:"environment"`
	PackageID   string `db:"packageId"`
}

// TableName returns the table name of PrivateEnvironment
func (p *PrivateEnvironment) TableName() string {
	return "privateenvironment"
}

// TableName returns the table name of EnvironmentVariable
func (ev *EnvironmentVariable) TableName() string {
	return "environmentvariable"
}

// PK defines the primary key of EnvironmentVariable
func (ev *EnvironmentVariable) PK() string {
	return "id"
}

// PK defines the primary key of PrivateEnvironment
func (p *PrivateEnvironment) PK() string {
	return "id"
}

// Save stores environmentvariable into database
func (ev *EnvironmentVariable) Save() error {
	ev.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(ev)
}

// Save stores private environement into database
func (p *PrivateEnvironment) Save() error {
	p.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(p)
}

// GetEnvironmentVariablesbyPackageID returns all environment variables in a single call
// Users are always requested to pass required repository/solver into the request
// TODO This can be improved in a future version, currently the db layer does
// TODO not support query all items with conditions.
func GetEnvironmentVariablesbyPackageID(packageID string, environment string) []EnvironmentVariable {
	environmentvariablesPointers := make([]*EnvironmentVariable, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&environmentvariablesPointers)
	var environmentVariables []EnvironmentVariable
	for i := range environmentvariablesPointers {
		if environment == "all" {
			if environmentvariablesPointers[i].PackageID == packageID {
				environmentVariables = append(environmentVariables, *environmentvariablesPointers[i])
			}
		} else {
			if environmentvariablesPointers[i].PackageID == packageID && environmentvariablesPointers[i].Environment == environment {
				environmentVariables = append(environmentVariables, *environmentvariablesPointers[i])
			}
		}
	}
	return environmentVariables
}

// MergeEnvironmentVariables merge all env vars into an array of strings
func MergeEnvironmentVariables(envs []EnvironmentVariable) []string {
	var envsStrings []string
	for i := range envs {
		envsStrings = append(envsStrings, envs[i].Key)
		envsStrings = append(envsStrings, envs[i].Value)
	}
	return envsStrings
}
