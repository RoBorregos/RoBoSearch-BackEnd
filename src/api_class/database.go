package api_class

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Code struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
	Filename string `bson:"filename" json:"filename"`
	Code string `bson:"code" json:"code"`
}

var Database *mongo.Database = nil

func InitConnection(url string) (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 7*time.Second)
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	Database = client.Database("apiclassdb")

	return Database, nil
}