// Package contrib defines extra stuff for cvpm
//Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*This file handles third party contributions and libraries.*/

package contrib

import (
	"github.com/BurntSushi/toml"
	"github.com/levigross/grequests"
	"log"
)

type dataset struct {
	Name  string `db:"Name"`
	Desc  string `db:"Desc"`
	Tags  string `db:"Tags"`
	Files string `db:"Files"`
	Link  string `db:"Link"`
}

type datasets struct {
	Datasets []dataset
}

func fetchRegistryContent(remoteURL string) *grequests.Response {
	resp, err := grequests.Get(remoteURL, &grequests.RequestOptions{
		Headers: map[string]string{"Used-By": "CVPM-Client"},
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp.Ok != true {
		log.Fatal("Bad Response from Index")
	}
	return resp
}

func addNewRegistry(remoteURL string) {
	// Add New Registry to local database
	var datasetsInRegistry datasets

	content := fetchRegistryContent(remoteURL).String()
	if _, err := toml.Decode(content, &datasetsInRegistry); err != nil {
		panic(err)
	}
	sess := GetDBInstance()
	datasetCollection := sess.Collection("dataset")
	for _, item := range datasetsInRegistry.Datasets {
		err := datasetCollection.InsertReturning(&item)
		if err != nil {
			panic(err)
		}
	}
	CloseDB(sess)
}

func fetchAllDatasets() []dataset {
	sess := GetDBInstance()
	datasetCollection := sess.Collection("dataset")
	res := datasetCollection.Find()
	var datasets []dataset

	err := res.All(&datasets)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}
	CloseDB(sess)
	return datasets
}
