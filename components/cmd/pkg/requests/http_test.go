package requests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	c := Client{Endpoint: ""}
	resp := c.get("http://httpbin.org/get", nil)
	assert.NotEmpty(t, resp.String())
}

func TestPost(t *testing.T) {
	c := Client{Endpoint: ""}
	resp := c.post("http://httpbin.org/post", map[string]string{"One": "Two"})
	assert.NotEmpty(t, resp.String())
}
