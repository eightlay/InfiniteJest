package main

import (
	"fmt"
	"runtime"

	"github.com/eightlay/InfiniteJest/iternal/graphics"
	"github.com/eightlay/InfiniteJest/iternal/graphics/shaders"
	"github.com/eightlay/InfiniteJest/iternal/graphics/textures"
	"github.com/eightlay/InfiniteJest/ui"
)

var (
	// Vertices of the triangle
	triangle = []float32{
		// positions     // texture coords
		+0.0, +0.5, 0.0, 0.0, 0.0, // top
		-0.5, -0.5, 0.0, 1.0, 0.0, // left
		+0.5, -0.5, 0.0, 0.6, 1.0, // right
	}
	// Triangle's vertices indices to use
	indices = []uint32{0, 1, 2}
)

func main() {
	// Lock thread
	runtime.LockOSThread()

	// Init graphics driver
	err := ui.Init(false)
	if err != nil {
		panic(fmt.Errorf("could not initilize graphics driver: %v", err))
	}
	defer ui.Terminate()

	// Create window
	window, err := ui.CreateWindow("Test", 500, 500, true)
	if err != nil {
		panic(fmt.Errorf("could not create window: %v", err))
	}
	defer window.Destroy()

	// Create shader
	shader, err := shaders.GetDefaultShader()
	if err != nil {
		panic(fmt.Errorf("could not create shader: %v", err))
	}

	// Create drawable
	drawable, err := graphics.CreateDrawable(triangle, indices)
	if err != nil {
		panic(fmt.Errorf("could not create drawable: %v", err))
	}

	// Create texture
	texture, err := textures.CreateTexture("wall.png")
	if err != nil {
		panic(fmt.Errorf("could not create texture: %v", err))
	}

	// Create camera
	camera, err := ui.CreateCamera([3]float32{0, 0, 0}, [3]float32{0, 1, 0}, -90, 0)
	if err != nil {
		panic(fmt.Errorf("could not create camera: %v", err))
	}

	// Handle window
	for window.IsRunning() {
		window.Clear()

		texture.Bind()
		shader.Use()

		view := camera.ViewMatrix()
		shader.SetMat4("view", &view[0])

		drawable.Draw()

		window.Handle()
	}
}
