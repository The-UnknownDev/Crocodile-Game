package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"bot/config"
)

var (
	ctx      = context.Background()
	database *mongo.Database
)

func Initialize() error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.C.Mongo.Uri))
	if err != nil {
		return err
	}
	database = client.Database(config.C.Mongo.Database)
	if err = usersInitialize(); err != nil {
		return err
	}
	if err = chatsInitialize(); err != nil {
		return err
	}
	return scoresInitialize()
}
