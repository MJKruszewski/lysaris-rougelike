package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
	"math/rand"
)

type Map struct {
	DrawInterface
}

var size = rl.Vector2{X: float32(config.GlobalConfig.TileSize), Y: float32(config.GlobalConfig.TileSize)}
var colors = []color.RGBA{
	color.RGBA{
		R: 27,
		G: 138,
		B: 39,
		A: 255,
	},
	color.RGBA{
		R: 138,
		G: 138,
		B: 39,
		A: 255,
	},
	color.RGBA{
		R: 27,
		G: 138,
		B: 138,
		A: 255,
	},
	color.RGBA{
		R: 27,
		G: 27,
		B: 39,
		A: 255,
	},
}
var textOffset = int32(config.GlobalConfig.TileSize/2 - 1)
var black = color.RGBA{
	R: 0,
	G: 0,
	B: 0,
	A: 255,
}

func (m *Map) Draw() {
	for row := 0; row < config.GlobalConfig.MapPanelDim.Width/config.GlobalConfig.TileSize; row++ {
		for col := 0; col < config.GlobalConfig.MapPanelDim.Height/config.GlobalConfig.TileSize; col++ {
			destX := ((float32(row) - 0) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.MapPanelDim.OffsetX)
			destY := ((float32(col) - 0) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.MapPanelDim.OffsetY)

			rl.DrawRectangleV(rl.Vector2{X: destX, Y: destY}, size, colors[rand.Intn(len(colors))])

			rl.DrawText(",", int32(destX)+textOffset, int32(destY)+textOffset, int32(config.GlobalConfig.TileFontSize), black)
		}
	}
}
