package gomongoengine

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ExampleStruct struct {
	Name string
}

func TestSave(t *testing.T) {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27027"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	t.Error(err)
	database := client.Database("testing")
	c := Client{Database: database}
	es := ExampleStruct{Name: "BOB"}
	res := c.Save(&es)
	t.Log(res)
	t.Error("todo")

}
