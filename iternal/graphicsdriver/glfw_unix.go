//go:build !windows
// +build !windows

package graphicsdriver

import "github.com/go-gl/glfw/v3.3/glfw"

func InitGLFW_OS() {
	// When compiling on Mac and using this specific version of OpenGL,
	// it’s a required hint for it to work.
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}
