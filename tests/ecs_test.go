package tests

import (
	"testing"

	"github.com/eightlay/InfiniteJest/ecs"
	ecs_example "github.com/eightlay/InfiniteJest/examples/ecs/ecs"
)

func TestECS(t *testing.T) {
	start := 5
	target := 4

	w := ecs.World{}
	w.Init()
	w.AddSystem("Physics", &ecs_example.Physics{})
	e := w.CreateEntity(&ecs_example.Position{Y: start})
	w.Update()

	c, ok := e.Component("Position")
	if !ok {
		panic("Wrong entity in system")
	}

	pos := c.(*ecs_example.Position)

	if pos.Y != target {
		t.Fatalf("System doesn't work")
	}
}
