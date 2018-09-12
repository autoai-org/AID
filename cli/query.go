package main

import (
	"log"
	"github.com/levigross/grequests"
)

func ClientPost(endpoint string, params map[string]string) {
	url := "http://127.0.0.1:10590/" + endpoint
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		JSON: params,
		IsAjax: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp.Ok != true {
		log.Fatal("Bad Response from Daemon")
	}
}
