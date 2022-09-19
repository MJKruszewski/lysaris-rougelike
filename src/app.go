package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gookit/event"
	"main/config"
	"main/consts"
	ui2 "main/ui"
	"main/ui/menu"
	"math/rand"
	"time"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(int32(config.GlobalConfig.WindowWidth), int32(config.GlobalConfig.WindowHeight), config.GlobalConfig.AppName)

	rl.SetTargetFPS(30)
	rl.SetWindowMinSize(800, 600)

	rand.Seed(time.Now().Unix())

	window := ui2.NewWindow()
	config.GameState.Window = &window

	window.AddPanel(&menu.MainMenu{})

	drawChannel := make(chan bool, 10)
	keyChannel := make(chan int32, 10)

	go eventLoop(drawChannel, keyChannel)
	renderLoop(&window, keyChannel, drawChannel)
}

func renderLoop(window *ui2.Window, keyChannel chan int32, drawChannel chan bool) {
	fmt.Println("Starting render loop")
	for !rl.WindowShouldClose() {
		resizeWindow()

		rl.BeginDrawing()

		for _, drawInterface := range window.GetPanels() {
			(*drawInterface).Draw()
		}

		rl.EndDrawing()

		keyPressed := rl.GetKeyPressed()
		if keyPressed > 0 {
			keyChannel <- keyPressed
			<-drawChannel
		}
	}
	fmt.Println("Ending render loop")

	rl.CloseWindow()
	close(keyChannel)
}

func eventLoop(out chan<- bool, in <-chan int32) {
	fmt.Println("Starting event loop")
	for {
		keyPressed, ok := <-in

		if !ok {
			fmt.Println("Stop event loop")

			break
		}

		fmt.Println("Pressed key: ", keyPressed)
		go func() {
			config.GameState.EventBus.MustFire(consts.KeyPressed, event.M{"key": keyPressed})
		}()
		out <- true
	}
	fmt.Println("Stop event loop")
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
