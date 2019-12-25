package requests

import (
	"github.com/levigross/grequests"
	"log"
)

// Client is the basic class for performing http requests
type Client struct {
	Endpoint string
}

func (c *Client) get(endpoint string, params map[string]string) *grequests.Response {
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

func (c *Client) post(endpoint string, params map[string]string) *grequests.Response {
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
