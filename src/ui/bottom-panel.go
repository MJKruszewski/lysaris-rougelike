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
	for row := 0; row < config.GlobalConfig.BottomPanelDim.Width/config.GlobalConfig.TileSize; row++ {
		for col := 0; col < config.GlobalConfig.BottomPanelDim.Height/config.GlobalConfig.TileSize; col++ {
			destX := ((float32(row) - 0) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.BottomPanelDim.OffsetX)
			destY := ((float32(col) - 0) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.BottomPanelDim.OffsetY)

			rl.DrawRectangleV(rl.Vector2{X: destX, Y: destY}, size, color.RGBA{
				R: 90,
				G: 0,
				B: 0,
				A: 255,
			})
		}
	}
}
