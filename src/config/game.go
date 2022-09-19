package config

import (
	"github.com/gookit/event"
	"main/game"
	"main/ui"
)

type Game struct {
	Player   *game.Player
	Map      *game.Map
	Window   *ui.Window
	EventBus *event.Manager
}

var GameState = NewGame()

func NewGame() Game {
	return Game{
		EventBus: event.NewManager("eventbus"),
	}
}
