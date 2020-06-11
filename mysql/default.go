package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func DB() *sqlx.DB {
	return db
}

func Init(conf Config) error {
	_Config = conf
	var err error
	db, err = New(conf)
	if err != nil {
		return err
	}
	return nil
}
