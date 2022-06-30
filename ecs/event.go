package ecs

// Event struct
//
// Event is synonym for runtime Entity creation
type Event struct {
	components []Component
}

func CreateEvent(comps ...Component) Event {
	return Event{comps}
}
