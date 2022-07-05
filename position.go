package infinitejest

import "github.com/eightlay/InfiniteJest/ecs"

type Position struct {
	X float64
	Y float64
}

func (Position) ID() ecs.ComponentID {
	return "Position"
}

func (p *Position) GetCoordinates() [2]float64 {
	return [2]float64{p.X, p.Y}
}

func (p *Position) GetX() float64 {
	return p.X
}

func (p *Position) GetY() float64 {
	return p.Y
}

func (p *Position) SetCoordinates(x, y float64) {
	p.X, p.Y = x, y
}

func (p *Position) SetX(x float64) {
	p.X = x
}

func (p *Position) SetY(y float64) {
	p.Y = y
}

func (p *Position) Add(x, y float64) {
	p.X += x
	p.Y += y
}

func (p *Position) AddX(d float64) {
	p.X += d
}

func (p *Position) AddY(d float64) {
	p.Y += d
}

func (p *Position) Sub(x, y float64) {
	p.X -= x
	p.Y -= y
}

func (p *Position) SubX(d float64) {
	p.X -= d
}

func (p *Position) SubY(d float64) {
	p.Y -= d
}

func (p *Position) Mult(x, y float64) {
	p.X *= x
	p.Y *= y
}

func (p *Position) MultX(m float64) {
	p.X *= m
}

func (p *Position) MultY(m float64) {
	p.Y *= m
}

func (p *Position) Div(x, y float64) {
	p.X /= x
	p.Y /= y
}

func (p *Position) DivX(d float64) {
	p.X /= d
}

func (p *Position) DivY(d float64) {
	p.Y /= d
}
