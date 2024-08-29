package storage

import (
	"context"
	"fmt"
	"hotel/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg config.Config)(*mongo.Collection, *mongo.Collection, *mongo.Collection, error){
	host:= cfg.Mongo_Host
	port:= cfg.Mongo_Port
	opt := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		return nil, nil, nil, err
	}
	col1 := client.Database(cfg.Mongo_Database).Collection(cfg.Hotel_Collectin)
	col2 := client.Database(cfg.Mongo_Database).Collection(cfg.Description_Collection)
	col3 := client.Database(cfg.Mongo_Database).Collection(cfg.Room_Collection)
	return col1, col2, col3, nil
}