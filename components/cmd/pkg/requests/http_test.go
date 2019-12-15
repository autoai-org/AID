package requests_test

import (
	"testing"
	"github.com/autoai-org/aiflow/components/aipm/pkg/requests"
)

func TestGet(t *testing.T) {
	c := requests.Client{Endpoint:""}
	resp := c.Get("http://httpbin.org/get", nil)
	if resp.String() == ""{
		t.Errorf("GET cannot fetch content")
	}
}

func TestPost(t *testing.T) {
	c := requests.Client{Endpoint:""}
	resp := c.Post("http://httpbin.org/post", map[string]string{"One": "Two"})
	if resp.String() == ""{
		t.Errorf("GET cannot fetch content")
	}
}