package storage

import (
	"context"
	"fmt"
	"user/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, cfg config.Config)(*mongo.Collection, error){
	host := cfg.Mongo_Host
	port := cfg.Mongo_Port
	opt := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}
	col := client.Database(cfg.Mongo_Database).Collection(cfg.Mongo_Collection)
	return col, nil
}