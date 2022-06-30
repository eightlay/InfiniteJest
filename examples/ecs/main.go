package main

import (
	"github.com/eightlay/InfiniteJest/ecs"
	ecs_example "github.com/eightlay/InfiniteJest/examples/ecs/ecs"
)

func main() {
	w := ecs.World{}
	w.Init()
	w.AddSystem("Physics", &ecs_example.Physics{})
	w.CreateEntity(&ecs_example.Position{Y: 5})
	for {
		w.Update()
	}
}
