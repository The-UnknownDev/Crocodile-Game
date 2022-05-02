package game

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/go-redis/redis/v8"
)

const GameDuration = 5 * time.Minute

var client *redis.Client

type Game struct {
	Chat int64
	Host int64
	Word int
}

func Initialize() error {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(client.Context()).Result()
	return err
}

func Poll(bot *gotgbot.Bot) {
	db := "0"
	ch := client.Subscribe(client.Context(), fmt.Sprintf("__keyevent@%s__:expired", db)).Channel()
	go func() {
		message := <-ch
		go func() {
			chatId, err := strconv.ParseInt(message.Payload, 10, 64)
			if err != nil {
				return
			}
			bot.SendMessage(chatId, "Time is up!", nil)
		}()
	}()
}

func (game *Game) Set() error {
	key := strconv.FormatInt(game.Chat, 10)
	value, err := json.Marshal(game)
	if err != nil {
		return err
	}
	_, err = client.Set(client.Context(), key, value, GameDuration).Result()
	return err
}

func (game *Game) Del() (bool, error) {
	key := strconv.FormatInt(game.Chat, 10)
	i, err := client.Del(client.Context(), key).Result()
	return i > 0, err
}

func Get(chat int64) (*Game, error) {
	game := &Game{Chat: chat}
	key := strconv.FormatInt(chat, 10)
	value, err := client.Get(client.Context(), key).Result()
	if err != nil {
		return game, err
	}
	err = json.Unmarshal([]byte(value), game)
	return game, err
}
