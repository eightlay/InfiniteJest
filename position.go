package infinitejest

import "github.com/eightlay/InfiniteJest/ecs"

// 2D position struct
type Position struct {
	X float64
	Y float64
}

// Position implements ecs.Component
func (Position) ID() ecs.ComponentID {
	return "Position"
}

// Get coordinates as float array{x, y}
func (p *Position) GetCoordinates() [2]float64 {
	return [2]float64{p.X, p.Y}
}

// Get X coordinate
func (p *Position) GetX() float64 {
	return p.X
}

// Get Y coordinate
func (p *Position) GetY() float64 {
	return p.Y
}

// Set coordinates
func (p *Position) SetCoordinates(x, y float64) {
	p.X, p.Y = x, y
}

// Set X coordinate
func (p *Position) SetX(x float64) {
	p.X = x
}

// Set Y coordinate
func (p *Position) SetY(y float64) {
	p.Y = y
}

// Add to coordinates
func (p *Position) Add(x, y float64) {
	p.X += x
	p.Y += y
}

// Add to X coordinate
func (p *Position) AddX(d float64) {
	p.X += d
}

// Add to Y coordinate
func (p *Position) AddY(d float64) {
	p.Y += d
}

// Substract from coordinates
func (p *Position) Sub(x, y float64) {
	p.X -= x
	p.Y -= y
}

// Substract from X coordinate
func (p *Position) SubX(d float64) {
	p.X -= d
}

// Substract from Y coordinate
func (p *Position) SubY(d float64) {
	p.Y -= d
}

// Multiply coordinates
func (p *Position) Mult(x, y float64) {
	p.X *= x
	p.Y *= y
}

// Multiply X coordinate
func (p *Position) MultX(m float64) {
	p.X *= m
}

// Multiply Y coordinate
func (p *Position) MultY(m float64) {
	p.Y *= m
}

// Divide coordinates
func (p *Position) Div(x, y float64) {
	p.X /= x
	p.Y /= y
}

// Divide X coordinate
func (p *Position) DivX(d float64) {
	p.X /= d
}

// Divide Y coordinate
func (p *Position) DivY(d float64) {
	p.Y /= d
}
