package redis

import (
	"testing"
)

func TestInitRedis(t *testing.T) {
	InitRedis([]string{"127.0.0.1:6379"}, "", 0, 50)
	result, err := defaultClient.Set("hello", "word", 0).Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
	val, err := defaultClient.Get("hello").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}
