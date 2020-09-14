package http_pool

//import (
//	"github.com/valyala/fasthttp"
//	"time"
//)
//
//var httpClient *fasthttp.Client
//
//func init() {
//	httpClient = &fasthttp.Client{
//		MaxConnsPerHost:     100,
//		ReadTimeout:         3 * time.Second,
//		MaxIdleConnDuration: 3 * time.Minute,
//		MaxConnWaitTimeout:  6 * time.Second,
//	}
//}
//
//func ClientForHttp() *fasthttp.Client {
//	return httpClient
//}
