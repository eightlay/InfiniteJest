package ecs

import "github.com/google/uuid"

// World struct
//
// World is high-level type. It implements the ECS itself.
// I.e. it operates with entities and systems.
type World struct {
	enm EntityManager
	sm  SystemManager
	evm EventManager
}

// Init world
func (w *World) Init() {
	w.enm.Init()
	w.sm.Init()
}

// Update world
func (w *World) Update() {
	w.sm.Update(&w.evm)
	w.ExecuteEvents()
}

// Execute world's events
func (w *World) ExecuteEvents() {
	for _, event := range w.evm.queue {
		w.CreateEntity(event.components...)
	}
	w.evm.Clear()
}

// Create new entity in the world
func (w *World) CreateEntity(comps ...Component) *Entity {
	e := w.enm.CreateEntity(comps...)
	w.sm.AddEntity(e)
	return e
}

// Delete world's entity by its id
func (w *World) DeleteEntity(id uuid.UUID) {
	w.sm.RemoveEntity(id)
	w.enm.DeleteEntity(id)
}

// Add system to the world
func (w *World) AddSystem(id string, s System) {
	w.sm.AddSystem(id, s)
}

// Delete system by its id
func (w *World) DeleteSystem(id string) {
	w.sm.DeleteSystem(id)
}

// Get all world's entities
func (w *World) Entities() []Entity {
	return w.enm.entities
}
