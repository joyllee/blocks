package http_pool

import (
	"github.com/valyala/fasthttp"
	"time"
)

var httpClient *fasthttp.Client

// max: Max connections per host
func InitClient(max int) {
	httpClient = &fasthttp.Client{
		MaxConnsPerHost:     max,
		ReadTimeout:         3 * time.Second,
		MaxIdleConnDuration: 3 * time.Minute,
		MaxConnWaitTimeout:  6 * time.Second,
	}
}

func Client() *fasthttp.Client {
	return httpClient
}
