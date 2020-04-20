// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package requests

import (
	"fmt"
	"log"

	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
	"github.com/levigross/grequests"
)

// Client is the basic class for performing http requests
type Client struct {
	Endpoint string
}

var defaultClient *Client

// NewClient returns a new empty client
func NewClient() *Client {
	if defaultClient != nil {
		return defaultClient
	}
	defaultClient = &Client{}
	return defaultClient
}

// Get sends a simple Get requests to specific endpoint
func (c *Client) Get(endpoint string, params map[string]string) *grequests.Response {
	url := c.Endpoint + endpoint
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Params: params,
	})
	if err != nil {
		utilities.CheckError(resp.Error, "Unable to make requests")
	}

	if resp.Ok != true {
		log.Println("Request did not return OK")
	}

	return resp
}

// Post sends a simple Post requests to specific endpoint
func (c *Client) Post(endpoint string, params map[string]string) *grequests.Response {
	url := c.Endpoint + endpoint
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:   params,
		IsAjax: true,
	})
	if err != nil {
		log.Println("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		log.Println(resp.String())
		log.Println("Request did not return OK")
	}
	return resp
}

// Upload sends a file into endpoint, via a post requests
func (c *Client) Upload(endpoint string, filename string) *grequests.Response {
	fmt.Printf("Uploading %s >>> %s", filename, endpoint)
	fd, err := grequests.FileUploadFromDisk(filename)
	if err != nil {
		utilities.CheckError(err, "Cannot open file: "+filename)
	}
	resp, err := grequests.Post(endpoint, &grequests.RequestOptions{
		Files: fd,
	})
	if err != nil {
		utilities.CheckError(err, "Cannot make requests to "+endpoint)
	}
	if resp.Ok != true {
		log.Println(resp.String())
		log.Println("Request did not return OK")
	}
	return resp
}
