package components

import (
	"github.com/eightlay/InfiniteJest/iternal/ecs"
	"github.com/eightlay/InfiniteJest/iternal/vector"
)

// Scale component
type Scale struct {
	scale vector.Vec3
}

// Create new scale
func CreateScale(x, y, z float32) (*Scale, error) {
	return &Scale{vector.Vec3{X: x, Y: y, Z: z}}, nil
}

// Scale implements ecs.Component
func (Scale) ID() ecs.ComponentID {
	return "Scale"
}

// Get coordinates as float array{x, y}
func (p *Scale) GetScale() *vector.Vec3 {
	return &p.scale
}
