package entites

type MovementType int

const (
	Walk MovementType = 0
	Fly  MovementType = 1
	Swim MovementType = 1
)

type Movable interface {
	IsAbleTo(movementType MovementType) bool
}
