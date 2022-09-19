package game

import (
	"main/game/entites"
)

type Player struct {
	X int
	Y int
}

func (p *Player) IsAbleTo(movementType entites.MovementType) bool {
	return movementType == entites.Walk
}
