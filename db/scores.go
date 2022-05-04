package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Score struct {
	ChatID int64 `bson:"chat_id"`
	UserID int64 `bson:"user_id"`
	Count  int64 `bson:"count"`
}

func scores() *mongo.Collection {
	return database.Collection("scores")
}

func scoresInitialize() error {
	_, err := scores().Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{"chat_id", 1}, {"user_id", 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	return err
}

func ScoresUpdate(chatId int64, userId int64) error {
	_, err := scores().UpdateOne(ctx, bson.D{{"chat_id", chatId}, {"user_id", userId}}, bson.D{{"$inc", bson.D{{"count", 1}}}}, options.Update().SetUpsert(true))
	return err
}

func ScoresTop(chatId int64) ([]Score, error) {
	cur, err := scores().Find(ctx, bson.D{{"chat_id", chatId}}, options.Find().SetSort(bson.D{{"count", -1}}).SetLimit(5))
	if err != nil {
		return nil, err
	}
	scores := []Score{}
	return scores, cur.All(ctx, &scores)
}
