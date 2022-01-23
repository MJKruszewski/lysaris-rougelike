package mapElements

import (
	"image/color"
	"main/game/entites"
	"math/rand"
)

type WaterTile struct {
	TileActions
	BaseTile
}

var waterColors = []color.RGBA{
	color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	},
}

var waterSymbols = []string{
	"^",
	"=",
}

func MakeWater(x int, y int) WaterTile {
	return WaterTile{
		BaseTile: BaseTile{
			X:      x,
			Y:      y,
			Symbol: waterSymbols[rand.Intn(len(waterSymbols))],
			Type:   Water,
			Color:  waterColors[rand.Intn(len(waterColors))],
		},
	}
}

func (t WaterTile) IsMovableTo(movable entites.Movable) bool {
	return movable.IsAbleTo(entites.Swim) || movable.IsAbleTo(entites.Fly)
}

func (t WaterTile) GetBaseTile() *BaseTile {
	return &t.BaseTile
}
