package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
	"main/game"
	"main/ui"
)

type Map struct {
	ui.DrawInterface
	CurrentMap *game.Map
	Player     *game.Player
	window     *ui.Window
}

var size = rl.Vector2{X: float32(config.GlobalConfig.TileSize), Y: float32(config.GlobalConfig.TileSize)}
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

	xSize := config.GlobalConfig.MapPanelDim.Width / config.GlobalConfig.TileSize
	ySize := config.GlobalConfig.MapPanelDim.Height / config.GlobalConfig.TileSize

	startRow := m.Player.X - (xSize / 2)
	startCol := m.Player.Y - (ySize / 2)

	for x := 0; x < xSize; x++ {
		startCol = m.Player.Y - (ySize / 2)
		for y := 0; y < ySize; y++ {

			if startRow < 0 || startCol < 0 || startRow >= len(m.CurrentMap.Tiles) || startCol >= len(m.CurrentMap.Tiles[startRow]) {
				startCol++
				continue
			}
			tile := m.CurrentMap.Tiles[startRow][startCol]

			destX := m.calculateXPos(x, 0)
			destY := m.calculateYPos(y, 0)

			rl.DrawRectangleV(rl.Vector2{X: destX, Y: destY}, size, tile.Color)
			rl.DrawText(tile.Symbol, int32(destX)+textOffset, int32(destY)+textOffset, int32(config.GlobalConfig.TileFontSize), black)

			startCol++
		}
		startRow++
	}

	rl.DrawText("@", int32(m.calculateXPos(xSize/2, 0))+textOffset, int32(m.calculateYPos(ySize/2, 0))+textOffset, int32(config.GlobalConfig.TileFontSize), black)
}

func (m *Map) calculateYPos(y int, startPosition int) float32 {
	return ((float32(y) - float32(startPosition)) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.MapPanelDim.OffsetY)
}

func (m *Map) calculateXPos(x int, startPosition int) float32 {
	return ((float32(x) - float32(startPosition)) * float32(config.GlobalConfig.TileSize)) + float32(config.GlobalConfig.MapPanelDim.OffsetX)
}

func (m *Map) SetWindow(window *ui.Window) {
	m.window = window
}
