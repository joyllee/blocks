package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var c *mongo.Client
var db *mongo.Database

func Client() *mongo.Client {
	return c
}

func DB() *mongo.Database {
	return db
}

func Init(conf Config) (err error) {
	_Config = conf
	c, db, err = New(conf)
	if err != nil {
		return err
	}
	return nil
}