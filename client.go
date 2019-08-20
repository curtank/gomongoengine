package gomongoengine

import (
	"context"
	"reflect"

	"github.com/curtank/gomongoengine/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Database *mongo.Database
}

func (c *Client) Save(ctx context.Context, o interface{}) interface{} {
	value := reflect.Indirect(reflect.ValueOf(o))
	collectionName := util.Caml2Snake(value.Type().Name())
	res, _ := c.Database.Collection(collectionName).InsertOne(ctx, o)

	return res
}
