package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joyllee/blocks/config"
	"github.com/joyllee/blocks/logger"
	"time"
)

const Nil = redis.Nil

var (
	defaultClient *redis.Client
)

func createRedisClient(address []string, password string, db, poolSize int) (*redis.Client, error) {
	if len(address) <= 0 {
		err := fmt.Errorf("single redis server is %s", address)
		return nil, err
	}
	redisClient := redis.NewClient(&redis.Options{
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
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

// InitRedis init redis connection and client
func InitRedis() (err error) {
	conf := config.ServerConfig.Redis
	address := conf.Addresses
	poolSize := conf.PoolSize
	db := conf.DB
	password := conf.Password

	defaultClient, err = createRedisClient(address, password, db, poolSize)
	if err != nil {
		logger.Error("redis connect error:", err)
		return
	}
	logger.Info("redis connect success")
	return err
}
