package main

import (
	ij "github.com/eightlay/InfiniteJest"

	"github.com/eightlay/InfiniteJest/ecs"
	ecs_example "github.com/eightlay/InfiniteJest/examples/ecs/ecs"
)

func main() {
	w := ecs.World{}
	w.Init()
	w.AddSystem("Physics", &ecs_example.Physics{})
	w.CreateEntity(&ij.Position{Y: 5})
	for {
		w.Update()
	}
}
