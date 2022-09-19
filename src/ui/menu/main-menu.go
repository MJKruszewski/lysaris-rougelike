package menu

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gookit/event"
	"image/color"
	"main/config"
	"main/consts"
	"main/ui"
)

type MainMenu struct {
	ui.DrawInterface
	window *ui.Window
}

func (m *MainMenu) GetName() string {
	return consts.MainMenu
}

func (m *MainMenu) Draw() {
	rl.DrawRectangleV(
		rl.Vector2{X: 0.0, Y: 0.0},
		rl.Vector2{X: float32(config.GlobalConfig.WindowWidth), Y: float32(config.GlobalConfig.WindowHeight)},
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		},
	)

	width := float32(100)
	height := float32(25)
	offsetX := width / 2

	for i, name := range [...]string{"New Game", "Continue", "Options", "Exit"} {
		pressed := rg.Button(rl.Rectangle{
			X:      float32(config.GlobalConfig.WindowWidth)/2 - offsetX,
			Y:      float32(config.GlobalConfig.WindowHeight)/2 + (height * float32(i)) - (height * 2),
			Width:  width,
			Height: height - 2,
		}, name)

		if pressed {
			config.GameState.EventBus.MustFire(consts.MainMenuEvent, event.M{"name": name})
		}
	}

}

func (m *MainMenu) SetWindow(window *ui.Window) {
	m.window = window
}

func (m *MainMenu) RegisterListeners() {
	config.GameState.EventBus.On(consts.MainMenuEvent, event.ListenerFunc(MainMenuItemSelectedSubscriber), event.Normal)
}

func (m *MainMenu) RemoveListeners() {
	config.GameState.EventBus.RemoveListeners(consts.MainMenuEvent)
}
