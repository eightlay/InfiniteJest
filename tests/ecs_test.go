package tests

import (
	"testing"

	ij "github.com/eightlay/InfiniteJest"

	"github.com/eightlay/InfiniteJest/ecs"
	ecs_example "github.com/eightlay/InfiniteJest/examples/ecs/ecs"
)

func TestECS(t *testing.T) {
	start := 5.0
	target := 4.0

	w := ecs.World{}
	w.Init()
	w.AddSystem("Physics", &ecs_example.Physics{})
	e := w.CreateEntity(&ij.Position{Y: start})
	w.Update()

	c, ok := e.Component("Position")
	if !ok {
		panic("Wrong entity in system")
	}

	pos := c.(*ij.Position)

	if pos.Y != target {
		t.Fatalf("System doesn't work")
	}
}
