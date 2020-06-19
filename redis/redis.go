package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const Nil = redis.Nil

var (
	defaultClient         *redis.Client
	defaultFailoverClient *redis.Client
	defaultClusterClient  *redis.ClusterClient
)

func InitDefault(conf Config) {
	if conf.ClientType == "cluster" {
		_, err := InitClusterRedis(conf)
		if err != nil {
			//logger.Fatal(err)
		}
		return
	}
	if conf.ClientType == "failover" {
		_, err := InitFailoverRedis(conf)
		if err != nil {
			//logger.Fatal(err)
		}
		return
	}
	_, err := InitRedis(conf)
	if err != nil {
		//logger.Fatal(err)
	}
}

// Client support failover or single mode client
func Client() *redis.Client {
	if defaultClient != nil {
		return defaultClient
	}
	if defaultFailoverClient != nil {
		return defaultFailoverClient
	}

	return nil
}

func FailoverClient() *redis.Client {
	return defaultFailoverClient
}

func ClusterClient() *redis.ClusterClient {
	return defaultClusterClient
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

func InitClusterRedis(conf Config) (*redis.ClusterClient, error) {
	if len(conf.Addresses) <= 0 {
		err := fmt.Errorf("cluster redis server are %s", conf.Addresses)
		return nil, err
	}
	defaultClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              conf.Addresses,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           conf.PoolSize,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: 500 * time.Millisecond,
	})
	_, err := defaultClusterClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return defaultClusterClient, nil
}

func InitFailoverRedis(conf Config) (*redis.Client, error) {
	if len(conf.Addresses) <= 0 {
		err := fmt.Errorf("cluster redis server are %s", conf.Addresses)
		return nil, err
	}
	defaultFailoverClient = redis.NewFailoverClient(&redis.FailoverOptions{
		SentinelAddrs:      conf.Addresses,
		MasterName:         conf.MasterName,
		DB:                 conf.DB,
		Password:           conf.Password,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           conf.PoolSize,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: 500 * time.Millisecond,
	})
	_, err := defaultFailoverClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return defaultFailoverClient, nil
}
