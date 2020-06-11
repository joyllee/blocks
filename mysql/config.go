package mysql

import "time"

type Config struct {
	DataSource
	MaxIdleConns int
	MaxOpenConns int
	KeepAlive    time.Duration
	//从库
	Slaves []DataSource
}

type DataSource struct {
	Driver   string
	Username string
	Password string
	Protocol string
	Address  string
	Port     string
	Dbname   string
	Params   string
}

var _Config Config

func Configuration() Config {
	return _Config
}
