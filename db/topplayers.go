package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopPlayer struct {
	Id        int64  `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Username  string `bson:"username"`
	Scores    int64  `bson:"scores"`
}

func TopPlayersGlobally() ([]TopPlayer, error) {
	cur, err := users().Aggregate(
		ctx,
		mongo.Pipeline{
			{
				{"$lookup", bson.D{
					{"from", "scores"},
					{"let", bson.D{
						{"id", "$id"},
					}},
					{"pipeline", bson.A{
						bson.D{
							{"$match", bson.D{
								{"$expr", bson.D{
									{"$eq", bson.A{"$user_id", "$$id"}},
								}},
							}},
						},
					}},
					{"as", "scores"},
				}},
			},
			{
				{"$set", bson.D{
					{"scores", bson.D{
						{"$sum", "$scores.count"},
					}},
				}},
			},
			{{"$sort", bson.D{
				{"scores", -1},
			}}},
			{{"$limit", 5}},
		},
	)
	if err != nil {
		return nil, err
	}
	topPlayers := []TopPlayer{}
	return topPlayers, cur.All(ctx, &topPlayers)
}

func TopPlayersInChat(chatId int64) ([]TopPlayer, error) {
	cur, err := scores().Aggregate(
		ctx,
		mongo.Pipeline{
			{
				{"$match", bson.D{
					{"$expr", bson.D{
						{"$eq", bson.A{"$chat_id", chatId}},
					}},
				}},
			},
			{
				{"$sort", bson.D{
					{"count", -1},
				}},
			},
			{
				{"$limit", 5},
			},
			{
				{"$set", bson.D{
					{"id", "$user_id"},
					{"scores", "$count"},
				}},
			},
			{
				{"$unset", bson.A{"chat_id", "user_id", "count"}},
			},
			{
				{"$lookup", bson.D{
					{"from", "users"},
					{"let", bson.D{
						{"id", "$id"},
					}},
					{"pipeline", bson.A{
						bson.D{
							{"$match", bson.D{
								{"$expr", bson.D{
									{"$eq", bson.A{"$id", "$$id"}},
								}},
							}},
						},
					}},
					{"as", "user"},
				}},
			},
			{
				{"$unwind", "$user"},
			},
			{
				{"$set", bson.D{
					{"first_name", "$user.first_name"},
					{"last_name", "$user.last_name"},
					{"username", "$user.username"},
				}},
			},
			{
				{"$unset", "user"},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	topPlayers := []TopPlayer{}
	return topPlayers, cur.All(ctx, &topPlayers)
}
