package ecs

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// System inteface.
//
// The user decides how the system stores entities.
//
// NeedComponents must return all required by the system components.
//
// It is strongly recommended to use IsSuitable function in
// AddEntity implementation.
type System interface {
	Init()
	Update(em *EventManager)
	NeedComponents() mapset.Set[ComponentID]
	AddEntity(e *Entity)
	RemoveEntity(id uuid.UUID)
}

// Check if given entity is suitable for given system
func IsSuitable(s System, e *Entity) bool {
	return s.NeedComponents().IsSubset(e.HasComponents())
}
