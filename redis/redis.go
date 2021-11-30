package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joyllee/blocks/logger"
	"time"
)

const Nil = redis.Nil

var (
	defaultClient *redis.Client
)

func InitDefault(conf Config) {
	_, err := InitRedis(conf)
	if err != nil {
		logger.Fatal(err)
	}
}

// InitRedis init redis connection and client
func InitRedis(conf Config) (*redis.Client, error) {
	if len(conf.Addresses) <= 0 {
		err := fmt.Errorf("single redis server is %s", conf.Addresses)
		return nil, err
	}
	defaultClient = redis.NewClient(&redis.Options{
		Addr:               conf.Addresses[0],
		Password:           conf.Password,
		DB:                 conf.DB,
		PoolSize:           conf.PoolSize,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: 500 * time.Millisecond,
	})
	_, err := defaultClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return defaultClient, nil
}
