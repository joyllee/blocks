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

func TestInit(t *testing.T) {
	err := Init(Config{
		URI:    "mongodb://127.0.0.1:27047",
		Dbname: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	one, err := DB().Collection("test").InsertOne(context.Background(), bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(one)

	one, err = DB().Collection("test").InsertOne(context.Background(), Person{
		Name: "jac",
		Age:  10,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(one)
}
