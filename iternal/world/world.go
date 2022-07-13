package world

import (
	"fmt"

	"github.com/eightlay/InfiniteJest/iternal/ecs"
	"github.com/eightlay/InfiniteJest/iternal/engine_ecs"
	"github.com/eightlay/InfiniteJest/iternal/graphics/render"
	"github.com/eightlay/InfiniteJest/iternal/graphics/shaders"
	"github.com/eightlay/InfiniteJest/iternal/graphicsdriver"
	"github.com/eightlay/InfiniteJest/iternal/ui"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/google/uuid"
)

// World struct
//
// World is high-level type. It implements the ECS itself.
// I.e. it operates with entities and systems.
type World struct {
	enm    ecs.EntityManager
	sm     ecs.SystemManager
	evm    ecs.EventManager
	Window *glfw.Window
}

// Init world
func (w *World) Init() {
	WINDOW_W := 800
	WINDOW_H := 600

	w.enm.Init()
	w.sm.Init()

	// Init graphics driver
	err := graphicsdriver.Init(true)
	if err != nil {
		panic(fmt.Errorf("could not initilize graphics driver: %v", err))
	}

	// Create window
	window, err := ui.CreateWindow("Test", WINDOW_W, WINDOW_H, true)
	if err != nil {
		panic(fmt.Errorf("could not create window: %v", err))
	}
	w.Window = window

	// Create camera
	camera, err := ui.CreateCamera([3]float32{0, 0, 3}, [3]float32{0, 1, 0}, -90, 0)
	if err != nil {
		panic(fmt.Errorf("could not create camera: %v", err))
	}

	// Create shader
	shader, err := shaders.GetDefaultShader()
	if err != nil {
		panic(fmt.Errorf("could not create shader: %v", err))
	}

	// Create engine's ECS config
	config, err := engine_ecs.CreateConfigEngineECS(shader, camera, WINDOW_W, WINDOW_H, window)
	if err != nil {
		panic(fmt.Errorf("could not create shader: %v", err))
	}

	// Renderer
	renderer := render.Renderer{}
	renderer.EngineSystemInit(config)
	w.AddSystem("Renderer", &renderer)
}

// Terminate world
func (w *World) Terminate() {
	w.Window.Destroy()
	graphicsdriver.Terminate()
}

// Update world
func (w *World) Update() {
	w.sm.Update(&w.evm)
	w.ExecuteEvents()
}

// Execute world's events
func (w *World) ExecuteEvents() {
	for _, event := range w.evm.Queue {
		w.CreateEntity(event.Components...)
	}
	w.evm.Clear()
}

// Create new entity in the world
func (w *World) CreateEntity(comps ...ecs.Component) *ecs.Entity {
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
func (w *World) AddSystem(id string, s ecs.System) {
	w.sm.AddSystem(id, s)
}

// Delete system by its id
func (w *World) DeleteSystem(id string) {
	w.sm.DeleteSystem(id)
}

// Get all world's entities
func (w *World) Entities() []ecs.Entity {
	return w.enm.Entities
}
