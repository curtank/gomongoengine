package gomongoengine

import (
	"context"
	"reflect"
	"time"

	"github.com/curtank/gomongoengine/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Database *mongo.Database
}

func (c *Client) Save(o interface{}) interface{} {
	value := reflect.Indirect(reflect.ValueOf(o))
	collectionName := util.Caml2Snake(value.Type().Name())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, _ := c.Database.Collection(collectionName).InsertOne(ctx, o)

	return res
}
