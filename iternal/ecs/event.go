package ecs

// Event struct
//
// Event is synonym for runtime Entity creation
type Event struct {
	Components []Component
}

func CreateEvent(comps ...Component) Event {
	return Event{comps}
}
