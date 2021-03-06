package menu

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
	"main/ui"
)

type MainMenu struct {
	ui.DrawInterface
	window *ui.Window
}

func (m *MainMenu) Draw() {
	rl.DrawRectangleV(
		rl.Vector2{X: float32(config.GlobalConfig.LeftPanelDim.OffsetX), Y: float32(config.GlobalConfig.LeftPanelDim.OffsetY)},
		rl.Vector2{X: float32(config.GlobalConfig.LeftPanelDim.Width), Y: float32(config.GlobalConfig.LeftPanelDim.Height)},
		color.RGBA{
			R: 0,
			G: 90,
			B: 0,
			A: 255,
		},
	)
}

func (m *MainMenu) SetWindow(window *ui.Window) {
	m.window = window
}
