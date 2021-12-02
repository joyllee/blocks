package config

type RedisConfig struct {
	Addresses []string `yaml:"addresses"`
	PoolSize  int      `yaml:"poolsize"`
	DB        int      `yaml:"db"`
	Password  string   `yaml:"password"`
}
