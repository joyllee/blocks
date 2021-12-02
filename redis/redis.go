package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const Nil = redis.Nil

var (
	defaultClient *redis.Client
)

// InitRedis init redis connection and client
func InitRedis(address []string, password string, db, poolSize int) (*redis.Client, error) {
	if len(address) <= 0 {
		err := fmt.Errorf("single redis server is %s", address)
		return nil, err
	}
	defaultClient = redis.NewClient(&redis.Options{
		Addr:               address[0],
		Password:           password,
		DB:                 db,
		PoolSize:           poolSize,
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
