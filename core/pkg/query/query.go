// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package query

import (
	"log"
	"os/user"

	"github.com/levigross/grequests"
	homedir "github.com/mitchellh/go-homedir"
)

// ClientPost triggers a POST Request to the server
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
	if !resp.Ok {
		log.Fatal("Bad Response from Daemon")
	}
	return resp
}

// ClientGet triggers a GET Request to the server
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
	if !resp.Ok {
		log.Fatal("Bad Response from Daemon")
	}
}

// ClientDelete triggers a DELETE Request to the server
func ClientDelete(endpoint string) *grequests.Response {
	url := "http://127.0.0.1:10590/" + endpoint
	resp, err := grequests.Delete(url, &grequests.RequestOptions{})
	if err != nil {
		log.Fatal(err)
	}
	if !resp.Ok {
		log.Fatal("Bad Response from Daemon")
	}
	return resp
}

// StopInferEngine stops python interface
func StopInferEngine(port string) bool {
	url := "http://127.0.0.1:" + port + "/"
	resp, err := grequests.Get(url, &grequests.RequestOptions{})
	if err != nil {
		log.Fatal(err)
	}
	if !resp.Ok {
		log.Fatal("Bad Response from Infer Engine")
		return false
	}
	return true
}
