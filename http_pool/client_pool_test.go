package http_pool

import (
	"github.com/valyala/fasthttp"
	"testing"
)

func TestClientPoolGetWithBody(t *testing.T) {
	InitClient(3)
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	req.Header.SetMethod(fasthttp.MethodGet)
	req.SetRequestURI("http://example.com")
	req.SetBodyString("test")

	err := Client().Do(req, res)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Body()) == 0 {
		t.Fatal("missing request body")
	}
	t.Log(string(res.Body()))
}
