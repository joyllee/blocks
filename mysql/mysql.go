package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joyllee/blocks/config"
	"github.com/joyllee/blocks/logger"
)

var DB *gorm.DB

func createDbClient(dbUrl string, maxIdleConns, maxOpenConns int) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		return
	}
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)
	return
}

func InitMysql() (err error) {
	conf := config.ServerConfig.Mysql.DbMysql
	dbUrl := conf.Url
	maxIdleConns := conf.MaxIdleConns
	maxOpenConns := conf.MaxOpenConns
	DB, err = createDbClient(dbUrl, maxIdleConns, maxOpenConns)
	DB.LogMode(true)
	if err != nil {
		logger.Error("mysql connect error:", err)
		return
	}
	logger.Info("mysql connect success")
	return err
}
