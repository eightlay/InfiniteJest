package event_example

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/eightlay/InfiniteJest/ecs"
	"github.com/google/uuid"
)

/* Poistion component */
type Position struct {
	Y int
}

func (*Position) ID() ecs.ComponentID {
	return "Position"
}

/* Physics system */
type Physics struct {
	entities map[uuid.UUID]*ecs.Entity
	gravity  int
}

func (p *Physics) Init() {
	p.entities = map[uuid.UUID]*ecs.Entity{}
	p.gravity = 1
}

func (p *Physics) Update(em *ecs.EventManager) {
	for _, e := range p.entities {
		c, ok := e.Component("Position")
		if !ok {
			panic("Wrong entity in system")
		}

		pos := c.(*Position)
		pos.Y -= p.gravity
	}
}

func (*Physics) NeedComponents() mapset.Set[ecs.ComponentID] {
	return mapset.NewSet(ecs.ComponentID("Position"))
}

func (p *Physics) AddEntity(e *ecs.Entity) {
	if ecs.IsSuitable(p, e) {
		p.entities[e.ID] = e
	}
}

func (p *Physics) RemoveEntity(id uuid.UUID) {
	delete(p.entities, id)
}

/* Bullet creation system */
type Bulleting struct {
	initY int
}

func (b *Bulleting) Init() {
	b.initY = 5
}

func (b *Bulleting) Update(em *ecs.EventManager) {
	em.Receive(ecs.CreateEvent(&Position{Y: b.initY}))
}

func (*Bulleting) NeedComponents() mapset.Set[ecs.ComponentID] {
	return mapset.NewSet[ecs.ComponentID]()
}

func (*Bulleting) AddEntity(e *ecs.Entity) {
}

func (*Bulleting) RemoveEntity(id uuid.UUID) {
}
