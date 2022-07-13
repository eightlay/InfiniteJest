package components

import (
	"github.com/eightlay/InfiniteJest/iternal/ecs"
	"github.com/eightlay/InfiniteJest/iternal/vector"
)

// Position component
type Position struct {
	pos vector.Vec3
}

// Create new position
func CreatePosition(x, y, z float32) (*Position, error) {
	return &Position{vector.Vec3{X: x, Y: y, Z: z}}, nil
}

// Position implements ecs.Component
func (Position) ID() ecs.ComponentID {
	return "Position"
}

// Get coordinates as float array{x, y}
func (p *Position) GetCoordinates() *vector.Vec3 {
	return &p.pos
}
