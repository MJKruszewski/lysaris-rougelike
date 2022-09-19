package menu

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gookit/event"
	"main/config"
	"main/consts"
	"main/game"
	game2 "main/ui/game"
)

func MainMenuItemSelectedSubscriber(e event.Event) error {
	keyPressed := e.Get("name").(string)
	fmt.Println("Event: " + keyPressed)

	switch keyPressed {
	case "New Game":
		mapper := game.Mapper{}
		gameMap := mapper.GenerateMap(100, 100)
		player := game.Player{
			X: len(gameMap.Tiles) / 2,
			Y: len(gameMap.Tiles[0]) / 2,
		}
		messageLog := game2.MessageLog{}

		config.GameState.Player = &player
		config.GameState.Map = &gameMap

		config.GameState.Window.AddPanel(&game2.LeftPanel{})
		config.GameState.Window.AddPanel(&game2.BottomPanel{
			MessageLog: &messageLog,
		})
		config.GameState.Window.AddPanel(&game2.Map{
			CurrentMap: &gameMap,
			Player:     &player,
		})
		config.GameState.Window.RemovePanel(consts.MainMenu)
	case "Continue":

	case "Options":

	case "Exit":
		rl.CloseWindow()
	}

	return nil
}
