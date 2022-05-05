package db

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id        int64  `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Username  string `bson:"username"`
}

func users() *mongo.Collection {
	return database.Collection("users")
}

func usersInitialize() error {
	_, err := users().Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{"id", 1}}, Options: options.Index().SetUnique(true)})
	return err
}

func UsersUpdate(user *gotgbot.User) error {
	_, err := users().UpdateOne(ctx, bson.D{{"id", user.Id}}, bson.D{{"$set", bson.D{{"first_name", user.FirstName}, {"last_name", user.LastName}, {"username", user.Username}}}}, options.Update().SetUpsert(true))
	return err
}

