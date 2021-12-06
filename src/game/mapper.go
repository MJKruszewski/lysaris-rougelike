package game

import (
	"image/color"
	"math/rand"
)

type Mapper struct {
}

type TileType int

const (
	Grass TileType = 0
	Water TileType = 1
)

type Tile struct {
	X      int
	Y      int
	Symbol string
	Type   TileType
	Color  color.RGBA
}

var colors = []color.RGBA{
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

var symbols = []string{
	",",
	"'",
	"\"",
}

func (m Mapper) GenerateMap(w int, h int) [][]Tile {
	result := make([][]Tile, w)

	for i := 0; i < w; i++ {
		result[i] = make([]Tile, h)
		for j := 0; j < h; j++ {
			result[i][j] = Tile{
				X:      i,
				Y:      j,
				Symbol: symbols[rand.Intn(len(symbols))],
				Type:   Grass,
				Color:  colors[rand.Intn(len(colors))],
			}
		}
	}

	return result
}
