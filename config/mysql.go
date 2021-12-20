package config

type MysqlConfig struct {
	DbMysql MysqlItemConfig `yaml:"db-mysql"`
}

type MysqlItemConfig struct {
	Url          string `yaml:"url"`            // 数据库地址
	MaxIdleConns int    `yaml:"max-idle-conns"` // 最大空闲数
	MaxOpenConns int    `yaml:"max-open-conns"` // 最大连接数
}
