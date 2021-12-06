package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
	"main/game"
)

type Map struct {
	DrawInterface
	CurrentMap [][]game.Tile
	Player     *game.Player
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
	rl.DrawRectangleV(
		rl.Vector2{X: float32(config.GlobalConfig.MapPanelDim.OffsetX), Y: float32(config.GlobalConfig.MapPanelDim.OffsetY)},
		rl.Vector2{X: float32(config.GlobalConfig.MapPanelDim.Width), Y: float32(config.GlobalConfig.MapPanelDim.Height)},
		black,
	)

	startRow := m.Player.X
	startCol := m.Player.Y

	for x := 0; x < config.GlobalConfig.MapPanelDim.Width/config.GlobalConfig.TileSize; x++ {
		startCol = m.Player.Y
		for y := 0; y < config.GlobalConfig.MapPanelDim.Height/config.GlobalConfig.TileSize; y++ {

			if startRow >= len(m.CurrentMap) || startCol >= len(m.CurrentMap[startRow]) {
				return
			}
			tile := m.CurrentMap[startRow][startCol]

			destX := m.calculateXPos(x, 0)
			destY := m.calculateYPos(y, 0)

			rl.DrawRectangleV(rl.Vector2{X: destX, Y: destY}, size, tile.Color)
			rl.DrawText(tile.Symbol, int32(destX)+textOffset, int32(destY)+textOffset, int32(config.GlobalConfig.TileFontSize), black)

			startCol++
		}
		startRow++
	}

	//todo adjust to allow move to corers
	rl.DrawText("@", int32(m.calculateXPos(config.GlobalConfig.MapPanelDim.Width/config.GlobalConfig.TileSize/2, 0))+textOffset, int32(m.calculateYPos(config.GlobalConfig.MapPanelDim.Height/config.GlobalConfig.TileSize/2, 0))+textOffset, int32(config.GlobalConfig.TileFontSize), black)
}

func (m *Map) calculateYPos(y int, startPosition int) float32 {
	return ((float32(y) - float32(startPosition)) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.MapPanelDim.OffsetY)
}

func (m *Map) calculateXPos(x int, startPosition int) float32 {
	return ((float32(x) - float32(startPosition)) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.MapPanelDim.OffsetX)
}

//  0 1 2 3 4 5 6
//0 x x x x x x x
//1 x x x x x x x
//2 x x x y x x x
//3 x x x x x x x
//4 x x x x x x x
