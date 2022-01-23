package mapElements

import (
	"image/color"
	"main/game/entites"
)

type BaseTile struct {
	Tile
	X      int
	Y      int
	Symbol string
	Type   TileType
	Color  color.RGBA
}

type Tile interface {
	GetX() int
	GetY() int
	GetSymbol() string
	GetType() TileType
	GetColor() color.RGBA
}

type TileActions interface {
	MovePossibility
	GetBaseTile() *BaseTile
}

type MovePossibility interface {
	IsMovableTo(movable entites.Movable) bool
}

type TileType int

const (
	Grass TileType = 0
	Water TileType = 1
)

func (t BaseTile) GetX() int {
	return t.X
}

func (t BaseTile) GetY() int {
	return t.Y
}

func (t BaseTile) GetSymbol() string {
	return t.Symbol
}

func (t BaseTile) GetType() TileType {
	return t.Type
}

func (t BaseTile) GetColor() color.RGBA {
	return t.Color
}
