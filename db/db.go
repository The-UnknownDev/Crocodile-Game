package db

import (
	"context"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()
var db *mongo.Database



type User struct {
	Id        int64  `bson:"id"`
	Scores    int64  `bson:"scores,omitempty"`
}

func Initialize() error {
	client, err := mongo.Connect(ctx)
	if err != nil {
		return err
	}
	db = client.Database("crocodilegame")
	return nil
}

func Score(user *gotgbot.User) {
	db.Collection("j").UpdateOne(ctx, bson.D{{"id", user.Id}}, bson.D{{"$inc", bson.D{{"scores", 1}}}})
}
