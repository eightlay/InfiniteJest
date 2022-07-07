package main

import (
	"fmt"
	"runtime"

	"github.com/eightlay/InfiniteJest/iternal/graphics"
	"github.com/eightlay/InfiniteJest/iternal/graphics/shaders3D"
	"github.com/eightlay/InfiniteJest/iternal/graphics/textures"
	"github.com/eightlay/InfiniteJest/ui"
	"github.com/go-gl/glfw/v3.3/glfw"
	glm "github.com/go-gl/mathgl/mgl32"
)

const (
	WINDOW_W int = 800
	WINDOW_H int = 600
)

var (
	cube = []float32{
		// Positions      // Texture coords
		// Front
		+0.5, +0.5, +0.5, 1.0, 1.0, // top right
		+0.5, -0.5, +0.5, 1.0, 0.0, // bottom right
		-0.5, -0.5, +0.5, 0.0, 0.0, // bottom left
		-0.5, +0.5, +0.5, 0.0, 1.0, // top left

		// Back
		+0.5, +0.5, -0.5, 1.0, 1.0, // top right
		+0.5, -0.5, -0.5, 1.0, 0.0, // bottom right
		-0.5, -0.5, -0.5, 0.0, 0.0, // bottom left
		-0.5, +0.5, -0.5, 0.0, 1.0, // top left

		// Top
		+0.5, +0.5, -0.5, 1.0, 0.0, // top right
		-0.5, +0.5, -0.5, 1.0, 1.0, // top left

		// Bottom
		+0.5, -0.5, -0.5, 1.0, 1.0, // bottom right
		-0.5, -0.5, -0.5, 0.0, 1.0, // bottom left
	}

	indices = []uint32{
		// Front
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle

		// Back
		4, 5, 7, // first triangle
		5, 6, 7, // second triangle

		// Top
		0, 3, 8, // first triangle
		3, 9, 8, // second triangle

		// Bottom
		1, 2, 10, // first triangle
		2, 11, 10, // second triangle
	}
)

func main() {
	// Lock thread
	runtime.LockOSThread()

	// Init graphics driver
	err := ui.Init(true)
	if err != nil {
		panic(fmt.Errorf("could not initilize graphics driver: %v", err))
	}
	defer ui.Terminate()

	// Create window
	window, err := ui.CreateWindow("Test", WINDOW_W, WINDOW_H, true)
	if err != nil {
		panic(fmt.Errorf("could not create window: %v", err))
	}
	defer window.Destroy()

	// Create shader
	shader, err := shaders3D.GetDefaultShader()
	if err != nil {
		panic(fmt.Errorf("could not create shader: %v", err))
	}

	// Create drawable
	drawable, err := graphics.CreateDrawable(cube, indices)
	if err != nil {
		panic(fmt.Errorf("could not create drawable: %v", err))
	}

	// Create texture
	texture, err := textures.CreateTexture("wall.png")
	if err != nil {
		panic(fmt.Errorf("could not create texture: %v", err))
	}

	// Use shader
	shader.Use()

	// Model
	model := glm.HomogRotate3DX(glm.DegToRad(-55))
	shader.SetMat4("model", &model[0])

	// View
	view := glm.Translate3D(0, 0, -5)
	shader.SetMat4("view", &view[0])

	// Projection
	projection := glm.Perspective(
		glm.DegToRad(45), float32(WINDOW_W)/float32(WINDOW_H), 0.1, 100.0,
	)
	shader.SetMat4("projection", &projection[0])

	// Configure global settings
	// gl.Enable(gl.DEPTH_TEST)
	// gl.DepthFunc(gl.LESS)

	// Handle window
	for window.IsRunning() {
		window.Clear()

		texture.Bind()
		shader.Use()

		model = glm.HomogRotate3DX(float32(glfw.GetTime()))
		model = glm.HomogRotate3DY(float32(glfw.GetTime())).Mul4(model)
		shader.SetMat4("model", &model[0])

		drawable.Draw()

		window.Handle()
	}
}
