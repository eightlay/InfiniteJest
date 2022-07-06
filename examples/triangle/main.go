package main

import (
	"fmt"
	"runtime"

	"github.com/eightlay/InfiniteJest/iternal/graphics"
	"github.com/eightlay/InfiniteJest/iternal/graphics/shaders"
	"github.com/eightlay/InfiniteJest/ui"
)

var (
	// Vertices of the triangle
	triangleCoords = []float32{
		0.0, 0.5, 0.0, // top
		-0.5, -0.5, 0.0, // left
		0.5, -0.5, 0.0, //right
	}
	// Triangle's vertices indices to use
	triangleIndices = []uint32{0, 1, 2}
	// Texture coordinates
	textureCoords = []float32{
		0.0, 0.0, // lower-left corner
		1.0, 0.0, // lower-right corner
		0.5, 1.0, // top-center corner
	}
)

func main() {
	// Lock thread
	runtime.LockOSThread()

	// Init graphics driver
	err := ui.Init()
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

	// Create program
	shader, err := shaders.GetDefaultShader()
	if err != nil {
		panic(fmt.Errorf("could not create program: %v", err))
	}

	// Create Drawable
	drawable, err := graphics.CreateDrawable(triangleCoords, triangleIndices)
	if err != nil {
		panic(fmt.Errorf("could not create vao: %v", err))
	}

	// Handle window
	for window.IsRunning() {
		window.Clear()

		shader.Use()
		drawable.Draw()

		window.Handle()
	}
}
