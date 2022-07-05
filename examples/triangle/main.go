package main

import (
	"fmt"
	"runtime"

	"github.com/eightlay/InfiniteJest/iternal/shaders"
	"github.com/eightlay/InfiniteJest/iternal/vertices"
	"github.com/eightlay/InfiniteJest/ui"
)

var triangle = []float32{
	-0.5, -0.5, 0.0, // left
	0.5, -0.5, 0.0, // right
	0.0, 0.5, 0.0, // top
}

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

	// Create VAO
	vao, err := vertices.CreateVAO(triangle)
	if err != nil {
		panic(fmt.Errorf("could not create vao: %v", err))
	}

	// Handle window
	for window.IsRunning() {
		window.Clear()

		program.Use()
		vao.Draw()

		window.Handle()
	}
}
