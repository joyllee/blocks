package redis

import (
	"testing"
)

func TestInitRedisCluster(t *testing.T) {
	InitClusterRedis(Config{
		Addresses: []string{"172.16.9.221:26379", "172.16.9.222:26379", "172.16.9.223:26379"},
		PoolSize:  10,
	})
	result, err := ClusterClient().Set("hello", "cluster", 0).Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
	val, err := ClusterClient().Get("hello").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}

func TestInitRedis(t *testing.T) {
	InitRedis(Config{
		Addresses: []string{"172.16.9.221:26379", "172.16.9.222:26379", "172.16.9.223:26379"},
		PoolSize:  10,
	})
	result, err := Client().Set("hello", "single", 0).Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
	val, err := Client().Get("hello").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}

func TestInitFailoverClient(t *testing.T) {
	InitFailoverRedis(Config{
		Addresses:  []string{"172.16.9.221:26379", "172.16.9.222:26379", "172.16.9.223:26379"},
		PoolSize:   10,
		ClientType: "failover",
		MasterName: "mymaster",
	})
	result, err := FailoverClient().Set("hello", "failover", 0).Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

	val, err := FailoverClient().Get("hello").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}
