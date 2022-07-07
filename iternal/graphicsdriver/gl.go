package graphicsdriver

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Initialize OpenGL
func InitGL() error {
	return gl.Init()
}

func EnableDepthTest() {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
}
