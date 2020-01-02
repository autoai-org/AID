// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package requests

import (
	"github.com/levigross/grequests"
	"log"
)

// Client is the basic class for performing http requests
type Client struct {
	Endpoint string
}

// Get sends a simple Get requests to specific endpoint
func (c *Client) Get(endpoint string, params map[string]string) *grequests.Response {
	url := c.Endpoint + endpoint
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Params: params,
	})
	if err != nil {
		log.Println("Unable to make request", resp.Error)
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
		log.Println("Request did not return OK")
	}
	return resp

}
