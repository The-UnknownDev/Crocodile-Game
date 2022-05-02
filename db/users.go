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
	Username  string `bson:"username"`
	Scores    int64  `bson:"scores"`
}

func users() *mongo.Collection {
	return database.Collection("users")
}

func usersInitialize() error {
	_, err := users().Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{"id", 1}},
			Options: options.Index().SetName("id").SetUnique(true),
		},
	)
	return err
}

func UsersUpdateUser(user *gotgbot.User) (*mongo.UpdateResult, error) {
	return users().UpdateOne(ctx, bson.D{{"id", user.Id}}, bson.D{{"$set", bson.D{{"first_name", user.FirstName}, {"username", user.Username}}}, {"$inc", bson.D{{"scores", 1}}}}, options.Update().SetUpsert(true))
}
