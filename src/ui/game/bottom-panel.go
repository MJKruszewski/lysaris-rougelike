package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
	"main/ui"
)

type BottomPanel struct {
	ui.DrawInterface
	window *ui.Window
}

func (m *BottomPanel) Draw() {
	rl.DrawRectangleV(
		rl.Vector2{X: float32(config.GlobalConfig.BottomPanelDim.OffsetX), Y: float32(config.GlobalConfig.BottomPanelDim.OffsetY)},
		rl.Vector2{X: float32(config.GlobalConfig.BottomPanelDim.Width), Y: float32(config.GlobalConfig.BottomPanelDim.Height)},
		color.RGBA{
			R: 90,
			G: 0,
			B: 0,
			A: 255,
		},
	)
}

func (m *BottomPanel) SetWindow(window *ui.Window) {
	m.window = window
}
