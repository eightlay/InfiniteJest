package main

import (
	"fmt"
	"runtime"

	"github.com/eightlay/InfiniteJest/iternal/graphics"
	"github.com/eightlay/InfiniteJest/iternal/shaders"
	"github.com/eightlay/InfiniteJest/ui"
)

var (
	// Vertices of the rectangle
	rectangle = []float32{
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
	}
	// Rectangle's vertices indices to use
	indices = []uint32{
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
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
	program, err := shaders.GetBasicProgram()
	if err != nil {
		panic(fmt.Errorf("could not create program: %v", err))
	}

	// Create Drawable
	drawable, err := graphics.CreateDrawable(rectangle, indices)
	if err != nil {
		panic(fmt.Errorf("could not create vao: %v", err))
	}

	// Handle window
	for window.IsRunning() {
		window.Clear()

		program.Use()
		drawable.Draw()

		window.Handle()
	}
}
