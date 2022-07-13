package ecs

// Component ID
type ComponentID string

// Component interface.
//
// Every component must have a unique user-defined ID
// in the component's world
type Component interface {
	// ID() must be implemented by value
	ID() ComponentID
}
