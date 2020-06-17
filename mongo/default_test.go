package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestMongoInit(t *testing.T) {
	m, err := Init(Config{
		URI:    "mongodb://127.0.0.1:27017",
		Dbname: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	one, err := m.Collection("test").InsertOne(context.Background(), bson.D{{"test", "mongo"}, {"hello", "world"}, {"pi", 3.14159}})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(one)

	one, err = m.Collection("test").InsertOne(context.Background(), Person{
		Name: "test",
		Age:  10,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(one)
}
