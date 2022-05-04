package session

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/go-redis/redis/v8"

	"bot/config"
)

var client *redis.Client

func Initialize() error {
	client = redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Address,
		Username: config.C.Redis.Username,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.Database,
	})
	_, err := client.Ping(client.Context()).Result()
	return err
}

func Poll(bot *gotgbot.Bot) {
	gamePoll(bot)
}
