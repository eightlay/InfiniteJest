package render

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/eightlay/InfiniteJest/iternal/components"
	"github.com/eightlay/InfiniteJest/iternal/ecs"
	"github.com/eightlay/InfiniteJest/iternal/engine_ecs"
	"github.com/eightlay/InfiniteJest/iternal/graphics/model"
	"github.com/eightlay/InfiniteJest/iternal/graphics/shaders"
	"github.com/eightlay/InfiniteJest/iternal/ui"
	"github.com/go-gl/glfw/v3.3/glfw"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

type Renderer struct {
	components mapset.Set[ecs.ComponentID]
	entities   map[uuid.UUID]*ecs.Entity
	shader     *shaders.Shader
	camera     *ui.Camera
	window     *glfw.Window
}

func (r *Renderer) Init() {
	r.components = mapset.NewSet(
		ecs.ComponentID("Model"), ecs.ComponentID("Position"), ecs.ComponentID("Scale"),
	)
	r.entities = map[uuid.UUID]*ecs.Entity{}
}

func (r *Renderer) EngineSystemInit(config *engine_ecs.ConfigEngineECS) {
	r.shader = config.Shader()
	r.camera = config.Camera()
	r.window = config.Window()
}

func (r *Renderer) Update(em *ecs.EventManager) {
	// Current window's parameters
	width_int, height_int := r.window.GetSize()
	centerX, centerY := float32(width_int)/2, float32(height_int)/2
	aspect := centerX / centerY

	ui.Clear()

	r.shader.Use()

	projection := glm.Perspective(
		r.camera.Zoom(), aspect, 0.1, 100.0,
	)
	r.shader.SetMat4("projection", &projection[0])

	view := r.camera.ViewMatrix()
	r.shader.SetMat4("view", &view[0])

	for _, entity := range r.entities {
		// Position
		p, ok := entity.Component("Position")
		if !ok {
			panic("Wrong entity in system")
		}
		position := p.(*components.Position).GetCoordinates()

		// Scale
		s, ok := entity.Component("Scale")
		if !ok {
			panic("Wrong entity in system")
		}
		scale := s.(*components.Scale).GetScale()

		trans := glm.Translate3D(position.X, position.Y, position.Z)
		trans = trans.Mul4(glm.Scale3D(scale.X, scale.Y, scale.Z))
		r.shader.SetMat4("model", &trans[0])

		// Draw Model
		m, ok := entity.Component("Model")
		if !ok {
			panic("Wrong entity in system")
		}
		model := m.(*model.Model)

		model.BindTexture()
		model.Draw()
	}

	ui.HandleWindow(r.window)
}

func (r *Renderer) NeedComponents() mapset.Set[ecs.ComponentID] {
	return r.components
}

func (r *Renderer) AddEntity(e *ecs.Entity) {
	if ecs.IsSuitable(r, e) {
		r.entities[e.ID] = e
	}
}

func (r *Renderer) RemoveEntity(id uuid.UUID) {
	delete(r.entities, id)
}
