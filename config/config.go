package config

import (
	"github.com/joyllee/blocks/kafka"
	"github.com/joyllee/blocks/mongo"
	"github.com/joyllee/blocks/mysql"
	"github.com/joyllee/blocks/redis"
)

var ServerConfig Config

type Config struct {
	Mode   string `default:"release"`
	Port   int32  `default:"62004"`
	Logger struct {
		LogLevel    string `default:"errors"`
		LogDir      string `default:"/opt/log"`
		LogFileName string `default:"demo.log"`
		LogFormat   string `default:"text"` // text or json
	}
	Kafka kafka.Config
	Mysql mysql.Config
	Mongo mongo.Config
	Redis redis.Config
}
