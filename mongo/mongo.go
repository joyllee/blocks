package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func New(conf Config) (c *mongo.Client, db *mongo.Database, err error) {
	c, err = mongo.NewClient(options.Client().ApplyURI(conf.URI))
	if err != nil {
		return nil, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = c.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}
	ctx, _ = context.WithTimeout(context.Background(), 20*time.Second)
	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}
	db = c.Database(conf.Dbname)
	return c, db, nil
}
