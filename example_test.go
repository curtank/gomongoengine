package gomongoengine

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//example struct A
type A struct {
	Name string
}

//TestConnection
func TestConnection(t *testing.T) {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27027"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	t.Error(err)
	collection := client.Database("testing").Collection("numbers")
	_, err = collection.InsertOne(ctx, A{Name: "sdklfjk"})
	t.Error(err)
	// time.Sleep(10 * time.Second)
	_, err = collection.InsertOne(ctx, A{Name: "sss"})
	t.Error(err)
	// time.Sleep(10 * time.Second)
	_, err = collection.InsertOne(ctx, A{Name: "qqq"})
	t.Error(err)
}
func TestType(t *testing.T) {
	type Empt struct {
		S     int32
		FN    string
		CCRE  A
		CCREP *A
	}
	e := Empt{S: 0}
	t.Log(e)
	t.Error("todo")
}
