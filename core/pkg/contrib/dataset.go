package contrib

import (
	"github.com/BurntSushi/toml"
	"github.com/cvpm-contrib/database"
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
	sess := database.GetDBInstance()
	datasetCollection := sess.Collection("dataset")
	for _, item := range datasetsInRegistry.Datasets {
		err := datasetCollection.InsertReturning(&item)
		if err != nil {
			panic(err)
		}
	}
	database.CloseDB(sess)
}

func fetchAllDatasets() []dataset {
	sess := database.GetDBInstance()
	datasetCollection := sess.Collection("dataset")
	res := datasetCollection.Find()
	var datasets []dataset

	err := res.All(&datasets)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}
	database.CloseDB(sess)
	return datasets
}
