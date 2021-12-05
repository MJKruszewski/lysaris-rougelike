package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"main/config"
	ui2 "main/ui"
	"math/rand"
	"time"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(int32(config.GlobalConfig.WindowWidth), int32(config.GlobalConfig.WindowHeight), "raylib [core] example - basic window")

	rl.SetTargetFPS(30)
	rl.SetWindowMinSize(800, 600)

	rand.Seed(time.Now().Unix())

	ui := []ui2.DrawInterface{
		&ui2.LeftPanel{},
		&ui2.BottomPanel{},
		&ui2.Map{},
	}

	for !rl.WindowShouldClose() {
		resizeWindow()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, drawInterface := range ui {
			drawInterface.Draw()
		}

		rl.EndDrawing()
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
