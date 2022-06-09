package redis

import (
	"github.com/joyllee/blocks/config"
	"github.com/joyllee/blocks/logger"
	"testing"
)

func TestInitRedis(t *testing.T) {
	logger.InitLogger()
	config.LoadConfig("../config/dev.yaml")
	err := InitRedis()
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
