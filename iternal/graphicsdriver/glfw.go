package graphicsdriver

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func InitGLFW() {
	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}
}

func TerminateGLFW() {
	glfw.Terminate()
}
