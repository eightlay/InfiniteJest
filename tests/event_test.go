package tests

import (
	"testing"

	"github.com/eightlay/InfiniteJest/ecs"
	event_example "github.com/eightlay/InfiniteJest/examples/events/events"
)

func TestEvents(t *testing.T) {
	target_len := 2
	target_pos := 4

	w := ecs.World{}
	w.Init()
	w.AddSystem("Bulleting", &event_example.Bulleting{})
	w.AddSystem("Physics", &event_example.Physics{})
	w.Update()
	w.Update()

	entities := w.Entities()

	if len(entities) != target_len {
		t.Fatalf("Bulleting doesn't work")
	}

	comp, _ := entities[0].Component("Position")
	pos := comp.(*event_example.Position)

	if pos.Y != target_pos {
		t.Fatalf("Physics doesn't work")
	}
}
