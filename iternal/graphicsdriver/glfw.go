package graphicsdriver

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Initialize GLFW
func InitGLFW() error {
	err := glfw.Init()

	// OpenGL context should be compatible with the major version 4
	// and minor version 1
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	// OpenGL profile in this case should be core.
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	// Depends on target OS
	InitGLFW_OS()

	return err
}

// Terminate GLFW
func TerminateGLFW() {
	glfw.Terminate()
}
