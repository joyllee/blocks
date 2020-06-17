package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	c  *mongo.Client
	db *mongo.Database
}

func (m *Mongo) Client() *mongo.Client {
	return m.c
}

func (m *Mongo) DB() *mongo.Database {
	return m.db
}

func Init(conf Config) (m *Mongo, err error) {
	_Config = conf
	c, db, err := New(conf)
	if err != nil {
		return
	}
	return &Mongo{
		c:  c,
		db: db,
	}, nil
}

func (m *Mongo) Collection(name string) *mongo.Collection {
	return m.DB().Collection(name)
}
