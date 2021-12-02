package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var ServerConfig Config

type Config struct {
	Mode   string `default:"release" yaml:"mode"`
	Port   int32  `default:"62004" yaml:"port"`
	Logger struct {
		LogLevel    string `default:"errors" yaml:"loglevel"`
		LogDir      string `default:"/opt/log" yaml:"logdir"`
		LogFileName string `default:"demo.log" yaml:"logfilename"`
		LogFormat   string `default:"text" yaml:"logformat"` // text or json
		LogWriter   bool   `default:"true" yaml:"logwriter"`
	} `yaml:"logger"`
	Redis RedisConfig `yaml:"redis"`
}

//LoadConfig 获取配置数据
func LoadConfig(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	err = yaml.Unmarshal(content, &ServerConfig)
	if err != nil {
		log.Fatalf("unmarshal config error: %v", err)
	}
}
