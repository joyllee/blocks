package redis

import (
	"testing"
)

func TestInitRedis(t *testing.T) {
	InitDefault(Config{
		Addresses: []string{"127.0.0.1:6379"},
		PoolSize:  10,
	})
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
