package ui

import (
	"fmt"

	"github.com/eightlay/InfiniteJest/graphics"
	"github.com/eightlay/InfiniteJest/iternal/graphicsdriver"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Window struct
type Window struct {
	window *glfw.Window
}

// Create a window with the given parameters
func CreateWindow(title string, width, height int, resizable bool) (*Window, error) {
	if !graphicsdriver.IsInitialized() {
		return nil, fmt.Errorf("you must call ui.Init() before calling CreateWindow")
	}

	// OpenGL context should be compatible with the major version 4
	// and minor version 1
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	// Is resizing window allowed
	if resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	// OpenGL profile in this case should be core.
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// When compiling on Mac and using this specific version of OpenGL,
	// it’s a required hint for it to work.
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}

	// MakeContextCurrent is what actually creates the OpenGL context
	// within the platform window that gets created by GLFW
	window.MakeContextCurrent()

	graphicsdriver.DefaultColor(graphics.CreateColor())

	return &Window{window}, nil
}

func (w *Window) IsRunning() bool {
	return !w.window.ShouldClose()
}

func (w *Window) Handle() {
	// Clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// Swap buffers, which swaps the front and back buffers of the window
	// (one used for rendering, the other for drawing)
	w.window.SwapBuffers()
	// Poll events, which actually allows us to catch events such as the close
	// button, so that the window can close (and, of course, all other events,
	// including input events)
	glfw.PollEvents()
}

func (w *Window) Destroy() {
	w.window.Destroy()
}
