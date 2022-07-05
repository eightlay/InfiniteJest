package graphicsdriver

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func InitGLFW() error {
	return glfw.Init()
}

func TerminateGLFW() {
	glfw.Terminate()
}
