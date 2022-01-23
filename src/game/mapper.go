package game

import (
	"main/game/mapElements"
	"math/rand"
)

type Mapper struct {
}

func (m *Mapper) GenerateMap(w int, h int) Map {
	result := make([][]mapElements.TileActions, w)

	for i := 0; i < w; i++ {
		result[i] = make([]mapElements.TileActions, h)
		for j := 0; j < h; j++ {
			result[i][j] = mapElements.MakeGrass(i, j)
		}
	}

	lakeStartX, lakeStartY := rand.Intn(len(result)), rand.Intn(len(result[0]))
	for i := 0; i < 5; i++ {
		if len(result) > lakeStartY+i {
			result[lakeStartX][lakeStartY+i] = mapElements.MakeWater(lakeStartX+i, lakeStartY+i)
		}

		if 0 <= lakeStartY-i {
			result[lakeStartX][lakeStartY-i] = mapElements.MakeWater(lakeStartX-i, lakeStartY-i)
		}
	}

	return Map{
		Tiles: result,
	}
}
