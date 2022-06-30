package main

import (
	"github.com/eightlay/InfiniteJest/ecs"
	event_example "github.com/eightlay/InfiniteJest/examples/events/events"
)

func main() {
	w := ecs.World{}
	w.Init()
	w.AddSystem("Physics", &event_example.Physics{})
	w.AddSystem("Bulleting", &event_example.Bulleting{})
	for {
		w.Update()
	}
}
