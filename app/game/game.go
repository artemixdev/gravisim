package game

import (
	"time"

	"gravisim/app/model"
)

type Game struct {
	lastUpdate time.Time
	bodies     []*model.Body
}

func New() *Game {
	return &Game{
		lastUpdate: time.Unix(0, 0),
		bodies:     make([]*model.Body, 0),
	}
}
