package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gookit/event"
	"main/config"
	"main/events"
	"main/game"
	"main/game/message"
	ui2 "main/ui"
	game2 "main/ui/game"
	"math/rand"
	"time"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(int32(config.GlobalConfig.WindowWidth), int32(config.GlobalConfig.WindowHeight), config.GlobalConfig.AppName)

	rl.SetTargetFPS(30)
	rl.SetWindowMinSize(800, 600)

	rand.Seed(time.Now().Unix())

	mapper := game.Mapper{}
	gameMap := mapper.GenerateMap(100, 100)
	player := game.Player{
		X: len(gameMap.Tiles) / 2,
		Y: len(gameMap.Tiles[0]) / 2,
	}
	messageLog := message.MessageLog{}

	window := ui2.NewWindow()
	window.AddPanel(&game2.LeftPanel{})
	window.AddPanel(&game2.BottomPanel{
		MessageLog: &messageLog,
	})
	window.AddPanel(&game2.Map{
		CurrentMap: &gameMap,
		Player:     &player,
	})

	drawChannel := make(chan bool, 10)
	keyChannel := make(chan int32, 10)

	event.On(events.KeyPressed, event.ListenerFunc(player.KeyPressedSubscriber), event.Normal)
	event.On(events.AddMessageLog, event.ListenerFunc(messageLog.AddMessageLogSubscriber), event.Normal)

	go eventLoop(drawChannel, keyChannel, &player, &gameMap)
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

func eventLoop(out chan<- bool, in <-chan int32, player *game.Player, gameMap *game.Map) {
	fmt.Println("Starting event loop")
	for {
		keyPressed, ok := <-in

		if !ok {
			fmt.Println("Stop event loop")

			break
		}

		fmt.Println("Pressed key: ", keyPressed)
		go func() {
			event.MustFire(events.KeyPressed, event.M{"key": keyPressed, "map": *gameMap})
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
