package ecs

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// Entity struct.
//
// Entity can have an unlimited number of unique components.
// Each entity has a unique auto-generated ID
type Entity struct {
	ID         uuid.UUID
	components map[ComponentID]Component
}

// Get IDs of each component in the entity
func (e *Entity) HasComponents() mapset.Set[ComponentID] {
	ids := mapset.NewSet[ComponentID]()

	for k := range e.components {
		ids.Add(k)
	}

	return ids
}

// Get entity's component.
// Returns false if the entity doesn't have the requested component
func (e *Entity) Component(id ComponentID) (Component, bool) {
	comp, ok := e.components[id]
	return comp, ok
}
