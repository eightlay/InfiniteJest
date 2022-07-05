package graphicsdriver

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func InitGL() error {
	return gl.Init()
}
