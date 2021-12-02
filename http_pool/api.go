package http_pool

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
	"runtime/debug"
	"time"
)

func Call(header map[string]string, uri, method string, req, resp interface{}) error {
	request, response := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func(t time.Time) {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
		if err := recover(); err != nil {
			//logger.Error(nil, "%v", err)
			debug.PrintStack()
		}
	}(time.Now())

	request.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	request.Header.Set("Pragma", "no-cache")
	request.Header.Set("Expires", "0")
	request.Header.Set("Access-Control-Allow-Origin", "*")
	request.Header.Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE")
	request.Header.Set("Access-Control-Allow-Headers", "Accept-Encoding, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, whiteList, signature,timestamp,appcode")

	for k, v := range header {
		request.Header.Set(k, v)
	}
	request.Header.SetMethod(method)
	request.SetRequestURI(uri)
	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	request.SetBody(reqBody)

	if err = Client().Do(request, response); err != nil {
		return err
	}
	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("Call %s get http status code: %d status: %s",
			uri, response.StatusCode, response.String())
	}
	if len(response.Body()) == 0 {
		return errors.New("missing request body")
	}
	if err = json.Unmarshal(response.Body(), &resp); err != nil {
		return err
	}

	return nil
}
