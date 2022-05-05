package session

import (
	"encoding/json"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/go-redis/redis/v8"

	"bot/config"
)

var client *redis.Client

func Initialize() error {
	client = redis.NewClient(&redis.Options{Addr: config.C.Redis.Address, Username: config.C.Redis.Username, Password: config.C.Redis.Password, DB: config.C.Redis.Database})
	_, err := client.Ping(client.Context()).Result()
	return err
}

func Poll(bot *gotgbot.Bot) {
	gamePoll(bot)
}

func Set(key string, v interface{}, expiration time.Duration) error {
	value, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = client.Set(client.Context(), key, value, expiration).Result()
	return err
}

func Get(key string, v interface{}) error {
	value, err := client.Get(client.Context(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(value), v)
}

func Del(key string) error {
	_, err := client.Del(client.Context(), key).Result()
	return err
}
