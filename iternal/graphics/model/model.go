package model

import (
	"github.com/eightlay/InfiniteJest/iternal/ecs"
	"github.com/eightlay/InfiniteJest/iternal/graphics/drawable"
	"github.com/eightlay/InfiniteJest/iternal/graphics/textures"
)

type Model struct {
	texture  *textures.Texture
	drawable *drawable.Drawable
}

func CreateModel(texture *textures.Texture, drawable *drawable.Drawable) (*Model, error) {
	return &Model{texture, drawable}, nil
}

func (Model) ID() ecs.ComponentID {
	return "Model"
}

func (m *Model) BindTexture() {
	m.texture.Bind()
}

func (m *Model) Draw() {
	m.drawable.Draw()
}
