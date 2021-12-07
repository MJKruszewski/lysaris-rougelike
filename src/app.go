package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gookit/event"
	"main/config"
	"main/events"
	"main/game"
	ui2 "main/ui"
	game2 "main/ui/game"
	"math/rand"
	"time"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(int32(config.GlobalConfig.WindowWidth), int32(config.GlobalConfig.WindowHeight), "raylib [core] example - basic window")

	rl.SetTargetFPS(30)
	rl.SetWindowMinSize(800, 600)

	rand.Seed(time.Now().Unix())

	mapper := game.Mapper{}
	gameMap := mapper.GenerateMap(100, 100)
	player := game.Player{
		X: len(gameMap.Tiles) / 2,
		Y: len(gameMap.Tiles[0]) / 2,
	}

	window := ui2.NewWindow()
	window.AddPanel(&game2.LeftPanel{})
	window.AddPanel(&game2.BottomPanel{})
	window.AddPanel(&game2.Map{
		CurrentMap: &gameMap,
		Player:     &player,
	})

	event.On(events.KeyPressed, event.ListenerFunc(player.KeyPressedSubscriber), event.Normal)

	for !rl.WindowShouldClose() {
		resizeWindow()

		rl.BeginDrawing()
		//rl.ClearBackground(rl.Black)

		for _, drawInterface := range window.GetPanels() {
			(*drawInterface).Draw()
		}

		rl.EndDrawing()

		keyPressed := rl.GetKeyPressed()
		for keyPressed > 0 {
			event.MustFire(events.KeyPressed, event.M{"key": keyPressed, "map": gameMap})
			keyPressed = rl.GetKeyPressed()
		}

	}

	rl.CloseWindow()
}

func resizeWindow() {
	if rl.IsWindowResized() {
		fmt.Println("Event | Resize window")

		width := rl.GetScreenWidth()
		height := rl.GetScreenHeight()

		if rl.GetScreenWidth()%config.GlobalConfig.TileSize != 0 {
			width = width - (width % config.GlobalConfig.TileSize) + config.GlobalConfig.TileSize
		}
		if rl.GetScreenHeight()%config.GlobalConfig.TileSize != 0 {
			height = height - (height % config.GlobalConfig.TileSize) + config.GlobalConfig.TileSize
		}

		rl.SetWindowSize(width, height)
		config.GlobalConfig.RecalculateScreen(width, height)
	}
}
