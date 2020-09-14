package consul

import (
	"net/http"
	"testing"
)

func TestConsulInit(t *testing.T)  {
	err := InitDefault(Config{
		Address:    "127.0.0.0:8500",
		Scheme:     "test",
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	http.ListenAndServe(":8888",nil)
}
