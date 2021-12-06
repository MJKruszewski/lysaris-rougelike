package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
)

type BottomPanel struct {
	DrawInterface
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
