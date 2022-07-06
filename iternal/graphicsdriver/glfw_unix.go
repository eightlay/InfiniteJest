//go:build !windows
// +build !windows

package graphicsdriver

import "github.com/go-gl/glfw/v3.3/glfw"

// Initialize GLFW for Unix
func InitGLFW_OS() {
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}
