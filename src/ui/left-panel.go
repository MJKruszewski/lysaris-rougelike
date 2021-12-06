package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
)

type LeftPanel struct {
	DrawInterface
}

func (m *LeftPanel) Draw() {
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
