package main

import (
	"fmt"
	"runtime"

	"github.com/eightlay/InfiniteJest/iternal/components"
	"github.com/eightlay/InfiniteJest/iternal/graphics/drawable"
	"github.com/eightlay/InfiniteJest/iternal/graphics/model"
	"github.com/eightlay/InfiniteJest/iternal/graphics/textures"
	"github.com/eightlay/InfiniteJest/iternal/world"
)

var (
	cube = []float32{
		// Positions      // Texture coords
		// Front
		+1.0, +1.0, +1.0, 1.0, 1.0, // top right
		+1.0, -1.0, +1.0, 1.0, 0.0, // bottom right
		-1.0, -1.0, +1.0, 0.0, 0.0, // bottom left
		-1.0, +1.0, +1.0, 0.0, 1.0, // top left

		// Back
		+1.0, +1.0, -1.0, 1.0, 1.0, // top right
		+1.0, -1.0, -1.0, 1.0, 0.0, // bottom right
		-1.0, -1.0, -1.0, 0.0, 0.0, // bottom left
		-1.0, +1.0, -1.0, 0.0, 1.0, // top left

		// Top
		+1.0, +1.0, -1.0, 1.0, 0.0, // top right
		-1.0, +1.0, -1.0, 1.0, 1.0, // top left

		// Bottom
		+1.0, -1.0, -1.0, 1.0, 1.0, // bottom right
		-1.0, -1.0, -1.0, 0.0, 1.0, // bottom left
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

	// New world
	w := world.World{}
	w.Init()
	defer w.Terminate()

	// Create drawable
	drawable, err := drawable.CreateDrawable(cube, indices)
	if err != nil {
		panic(fmt.Errorf("could not create drawable: %v", err))
	}

	// Create texture
	texture, err := textures.CreateTexture("wall.png")
	if err != nil {
		panic(fmt.Errorf("could not create texture: %v", err))
	}

	// New entity
	position, err := components.CreatePosition(0, 0, 0)
	if err != nil {
		panic(fmt.Errorf("could not create position: %v", err))
	}

	scale, err := components.CreateScale(0.5, 0.5, 0.5)
	if err != nil {
		panic(fmt.Errorf("could not create scale: %v", err))
	}

	model, err := model.CreateModel(texture, drawable)
	if err != nil {
		panic(fmt.Errorf("could not create model: %v", err))
	}

	w.CreateEntity(position, scale, model)

	for !w.Window.ShouldClose() {
		w.Update()
	}
}
