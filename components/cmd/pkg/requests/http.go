package requests

import (
	"log"
	"github.com/levigross/grequests"
)

type Client struct {
	Endpoint string
}

func (c *Client) Get (endpoint string, params map[string]string) *grequests.Response {
	url := c.Endpoint + endpoint
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Params:  params,
	})
	if err != nil {
		log.Println("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		log.Println("Request did not return OK")
	}

	return resp
}

func (c *Client) Post (endpoint string, params map[string]string) *grequests.Response {
	url := c.Endpoint + endpoint
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:    params,
		IsAjax:  true,
	})
	if err != nil {
		log.Println("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		log.Println("Request did not return OK")
	}
	return resp

}