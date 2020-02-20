// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

// PrivateEnvironment allows users to add new environments other than dev, test, prod
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
