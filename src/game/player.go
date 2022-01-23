package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gookit/event"
	"main/game/entites"
)

type Player struct {
	X int
	Y int
}

func (p *Player) KeyPressedSubscriber(e event.Event) error {
	localMap := e.Get("map").(Map)
	keyPressed := e.Get("key").(int32)

	switch keyPressed {
	case rl.KeyUp:
		if p.Y == 0 || !localMap.Tiles[p.X][p.Y-1].IsMovableTo(p) {
			break
		}

		p.Y--
	case rl.KeyDown:
		if p.Y >= len(localMap.Tiles[0])-1 || !localMap.Tiles[p.X][p.Y+1].IsMovableTo(p) {
			break
		}

		p.Y++
	case rl.KeyLeft:
		if p.X == 0 || !localMap.Tiles[p.X-1][p.Y].IsMovableTo(p) {
			break
		}

		p.X--
	case rl.KeyRight:
		if p.X >= len(localMap.Tiles)-1 || !localMap.Tiles[p.X+1][p.Y].IsMovableTo(p) {
			break
		}

		p.X++
	}

	return nil
}

func (p *Player) IsAbleTo(movementType entites.MovementType) bool {
	return movementType == entites.Walk
}
