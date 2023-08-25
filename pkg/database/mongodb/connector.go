package mongodb

import (
	"context"
	"time"

	"github.com/pondparinya/go-template/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBClient(cfg config.Config) (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if cfg.MongoDB.EnableURI {
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI))
		return client, err
	}
	return client, nil
}
