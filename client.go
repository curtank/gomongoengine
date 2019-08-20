package gomongoengine

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Database *mongo.Database
}

func (c *Client) Save(o interface{}) {

}
