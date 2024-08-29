package storage

import (
	"booking/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg *config.Config, ctx context.Context)(*mongo.Collection, *mongo.Collection, error){
	host := cfg.Mongo_Host
	port := cfg.Mongo_Port
	opt := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s",host, port))
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil,nil, err
	}
	col := client.Database(cfg.Mongo_Database).Collection(cfg.Mongo_Collection)
	col2 := client.Database(cfg.Mongo_User_Database).Collection(cfg.Mongo_User_collection)
	return col, col2, nil
}