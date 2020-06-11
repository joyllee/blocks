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
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetConnMaxLifetime(conf.KeepAlive)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
