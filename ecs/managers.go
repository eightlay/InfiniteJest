package ecs

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

// EntityManager operates with entities
type EntityManager struct {
	entities []Entity
}

// Init EntityManager
func (sm *EntityManager) Init() {
	sm.entities = []Entity{}
}

// Create a new entity with the given components
func (em *EntityManager) CreateEntity(comps ...Component) *Entity {
	e := Entity{uuid.New(), map[ComponentID]Component{}}

	for _, c := range comps {
		id := c.ID()
		e.components[id] = c
	}

	em.entities = append(em.entities, e)

	return &em.entities[len(em.entities)-1]
}

// Delete the entity by its ID
func (em *EntityManager) DeleteEntity(id uuid.UUID) {
	i := slices.IndexFunc(em.entities, func(e Entity) bool { return e.ID == id })
	if i != -1 {
		em.entities = slices.Delete(em.entities, i, i+1)
	}
}

// SystemManager operates with systems
type SystemManager struct {
	systems map[string]System
}

// Init SystemManager
func (sm *SystemManager) Init() {
	sm.systems = map[string]System{}
}

// Call each system's Update method
func (sm *SystemManager) Update(em *EventManager) {
	for _, s := range sm.systems {
		s.Update(em)
	}
}

// Add the system with the given unique ID
func (sm *SystemManager) AddSystem(id string, s System) {
	s.Init()
	sm.systems[id] = s
}

// Delete the system by its id
func (sm *SystemManager) DeleteSystem(id string) {
	delete(sm.systems, id)
}

// Add the entity to all suitable sistems
func (sm *SystemManager) AddEntity(e *Entity) {
	for _, s := range sm.systems {
		s.AddEntity(e)
	}
}

// Remove the entity by its id from all the systems
// where the entity is presented
func (sm *SystemManager) RemoveEntity(id uuid.UUID) {
	for _, s := range sm.systems {
		s.RemoveEntity(id)
	}
}

// EventManager operates with events
//
// Events are executed in FIFO order
type EventManager struct {
	queue []Event
}

func (em *EventManager) Init() {
	em.queue = []Event{}
}

func (em *EventManager) Receive(e Event) {
	em.queue = append(em.queue, e)
}

func (em *EventManager) Clear() {
	em.queue = nil
}
