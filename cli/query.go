package main

import (
	"github.com/levigross/grequests"
	"github.com/mitchellh/go-homedir"
	"log"
	"os/user"
)

func ClientPost(endpoint string, params map[string]string) *grequests.Response {
	url := "http://127.0.0.1:10590/" + endpoint
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		Headers: map[string]string{"X-Current-User": currentUser.Username},
		JSON:    params,
		IsAjax:  true,
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp.Ok != true {
		log.Fatal("Bad Response from Daemon")
	}
	return resp
}

func ClientGet(endpoint string, params map[string]string) {
	url := "http://127.0.0.1:10590/" + endpoint
	myhomedir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Headers: map[string]string{"X-Current-Homedir": myhomedir},
		Params:  params,
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp.Ok != true {
		log.Fatal("Bad Response from Daemon")
	}
}
