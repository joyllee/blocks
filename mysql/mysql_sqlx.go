package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(conf Config) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", conf.Username, conf.Password, conf.Protocol, conf.Address, conf.Port, conf.Dbname, conf.Params)
	var err error
	db, err = sqlx.Connect("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(conf.MaxIdleConns) //设置连接池中的保持连接的最大连接数
	db.SetMaxOpenConns(conf.MaxOpenConns) //设置打开数据库的最大连接数
	db.SetConnMaxLifetime(conf.KeepAlive) //设置连接可以被使用的最长有效时间，如果过期，连接将被拒绝
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
