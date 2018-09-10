package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var store *leveldb.DB
var cacheInited bool

func initCache() {
	if cacheInited {
	} else {
		var err error
		store, err = leveldb.OpenFile(".data.db", nil)
		if err != nil {
			panic(err)
		}
		cacheInited = true
	}
}

func setCache(key string, value string) {
	initCache()
	if err := store.Put([]byte(key), []byte(value), nil); err != nil {
		log.Fatal(err)
	}
}

func getCache(key string) string {
	initCache()
	result, _ := store.Get([]byte(key), nil)
	return string(result)
}

func deleteKey(key string) *error {
	initCache()
	err := store.Delete([]byte(key), nil)
	return &err
}

func closeCache() {
	err := store.Close()
	if err != nil {
		panic(err)
	}
}
