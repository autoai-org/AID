// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/ilibs/gosql/v2"
	_ "github.com/mattn/go-sqlite3"
)

func save() {
	db := storage.NewDB("sqlite3", "./db.db")
	db.Connect()
	db.CreateTables()
	packages := &entities.Package{}
	gosql.Model(packages).Where("id=?",1).Get()
}