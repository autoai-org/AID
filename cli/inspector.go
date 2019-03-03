package main

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Total       int           `json:"total"`
	TotalPage   int           `json:"total_page"`
	CurrentPage int           `json:"current_page"`
	PerPage     int           `json:"per_page"`
	HasNext     bool          `json:"has_next"`
	HasPrev     bool          `json:"has_prev"`
	NextPageUrl string        `json:"next_page_url"`
	PrevPageUrl string        `json:"prev_page_url"`
	Data        []RequestStat `json:"data"`
}

type RequestStat struct {
	RequestedAt   time.Time   `json:"requested_at"`
	RequestUrl    string      `json:"request_url"`
	HttpMethod    string      `json:"http_method"`
	HttpStatus    int         `json:"http_status"`
	ContentType   string      `json:"content_type"`
	GetParams     interface{} `json:"get_params"`
	PostParams    interface{} `json:"post_params"`
	PostMultipart interface{} `json:"post_multipart"`
	ClientIP      string      `json:"client_ip"`
	Cookies       interface{} `json:"cookies"`
	Headers       interface{} `json:"headers"`
}

type AllRequests struct {
	Requests []RequestStat `json:"requests"`
}

var allRequests = AllRequests{}

func GetPaginator() AllRequests {
	return allRequests
}

func InspectorStats() gin.HandlerFunc {
	return func(c *gin.Context) {

		urlPath := c.Request.URL.Path

		if urlPath == "/_inspector" {

		} else {
			if strings.HasPrefix(urlPath, "/engine") {
				start := time.Now()
				c.Request.ParseForm()
				c.Request.ParseMultipartForm(10000)
				request := RequestStat{
					RequestedAt:   start,
					RequestUrl:    urlPath,
					HttpMethod:    c.Request.Method,
					HttpStatus:    c.Writer.Status(),
					ContentType:   c.ContentType(),
					Headers:       c.Request.Header,
					Cookies:       c.Request.Cookies(),
					GetParams:     c.Request.URL.Query(),
					PostParams:    c.Request.PostForm,
					PostMultipart: c.Request.MultipartForm,
					ClientIP:      c.ClientIP(),
				}
				allRequests.Requests = append([]RequestStat{request}, allRequests.Requests...)
			}
		}

		c.Next()

	}
}

func paginate(s []RequestStat, offset, length int) []RequestStat {
	end := offset + length
	if end < len(s) {
		return s[offset:end]
	}

	return s[offset:]
}
