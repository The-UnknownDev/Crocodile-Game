package db

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chat struct {
	Id       int64           `bson:"id"`
	Title    int64           `bson:"title"`
	Games    int64           `bson:"games"`
	UserWins map[int64]int64 `bson:"user_wins"`
}

func chats() *mongo.Collection {
	return database.Collection("chats")
}

func chatsInitialize() error {
	_, err := chats().Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{"id", 1}}, Options: options.Index().SetUnique(true)})
	return err
}

func ChatsUpdate(chat *gotgbot.Chat) error {
	_, err := chats().UpdateOne(ctx, bson.D{{"id", chat.Id}}, bson.D{{"$set", bson.D{{"title", chat.Title}}}, {"$inc", bson.D{{"games", 1}}}})
	return err
}
