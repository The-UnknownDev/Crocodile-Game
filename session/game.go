package session

import (
	"time"
)

const GameDuration = 5 * time.Minute

type Game struct {
	Host int64
	Word int
}
