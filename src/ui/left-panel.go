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
	for row := 0; row < config.GlobalConfig.LeftPanelDim.Width/config.GlobalConfig.TileSize; row++ {
		for col := 0; col < config.GlobalConfig.LeftPanelDim.Height/config.GlobalConfig.TileSize; col++ {
			destX := ((float32(row) - 0) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.LeftPanelDim.OffsetX)
			destY := ((float32(col) - 0) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.LeftPanelDim.OffsetY)

			rl.DrawRectangleV(rl.Vector2{X: destX, Y: destY}, size, color.RGBA{
				R: 0,
				G: 90,
				B: 0,
				A: 255,
			})
		}
	}
}
