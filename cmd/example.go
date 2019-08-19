package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//example struct A
type A struct {
	Name string
}

//TestConnection
func TestConnection() {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27027").SetServerSelectionTimeout(2 * time.Second))
	ctx := context.Background()
	err := client.Connect(ctx)
	collection := client.Database("testing").Collection("numbers")

	fmt.Println(err)
	for {
		_, err = collection.InsertOne(ctx, A{Name: "sdklfjk"})
		fmt.Println(err)
		time.Sleep(2 * time.Second)
	}
}
func main() {
	TestConnection()
}
