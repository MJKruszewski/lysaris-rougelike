package mapElements

import (
	"image/color"
	"main/game/entites"
	"math/rand"
)

type GrassTile struct {
	TileActions
	BaseTile
}

var grassColors = []color.RGBA{
	color.RGBA{
		R: 8,
		G: 182,
		B: 13,
		A: 255,
	},
	color.RGBA{
		R: 56,
		G: 229,
		B: 62,
		A: 255,
	},
	color.RGBA{
		R: 35,
		G: 187,
		B: 40,
		A: 255,
	},
	color.RGBA{
		R: 10,
		G: 211,
		B: 16,
		A: 255,
	},
}

var grassSymbols = []string{
	",",
	"'",
	"\"",
}

func MakeGrass(x int, y int) GrassTile {
	return GrassTile{
		BaseTile: BaseTile{
			X:      x,
			Y:      y,
			Symbol: grassSymbols[rand.Intn(len(grassSymbols))],
			Type:   Grass,
			Color:  grassColors[rand.Intn(len(grassColors))],
		},
	}
}

func (t GrassTile) IsMovableTo(movable entites.Movable) bool {
	return movable.IsAbleTo(entites.Walk)
}

func (t GrassTile) GetBaseTile() *BaseTile {
	return &t.BaseTile
}
