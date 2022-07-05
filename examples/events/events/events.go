package event_example

import (
	ij "github.com/eightlay/InfiniteJest"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/eightlay/InfiniteJest/ecs"
	"github.com/google/uuid"
)

/* Physics system */
type Physics struct {
	entities map[uuid.UUID]*ecs.Entity
	gravity  float64
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

		pos := c.(*ij.Position)
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
	initY float64
}

func (b *Bulleting) Init() {
	b.initY = 5.0
}

func (b *Bulleting) Update(em *ecs.EventManager) {
	em.Receive(ecs.CreateEvent(&ij.Position{Y: b.initY}))
}

func (*Bulleting) NeedComponents() mapset.Set[ecs.ComponentID] {
	return mapset.NewSet[ecs.ComponentID]()
}

func (*Bulleting) AddEntity(e *ecs.Entity) {
}

func (*Bulleting) RemoveEntity(id uuid.UUID) {
}
