package redis

import (
	"github.com/go-redis/redis/v8"

	"bot/config"
)

var Client *redis.Client

func Initialize() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Address,
		Username: config.C.Redis.Username,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.Database,
	})
	_, err := Client.Ping(Client.Context()).Result()
	return err
}
