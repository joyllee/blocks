package mysql

import "testing"

func TestInit(t *testing.T) {
	Init(Config{
		DataSource:   DataSource{},
		MaxIdleConns: 0,
		MaxOpenConns: 0,
		KeepAlive:    0,
		Slaves:       nil,
	})
}
