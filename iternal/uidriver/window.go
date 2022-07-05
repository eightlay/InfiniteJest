package uidriver

import (
	"fmt"

	"github.com/eightlay/InfiniteJest/iternal/graphicsdriver"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Create a window with the given parameters
func CreateWindow(title string, width, height int, resizable bool) (*glfw.Window, error) {
	if !graphicsdriver.IsInitialized() {
		return nil, fmt.Errorf("you must call ui.Init() before calling CreateWindow")
	}

	// Is resizing window allowed
	if resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}

	// MakeContextCurrent is what actually creates the OpenGL context
	// within the platform window that gets created by GLFW
	window.MakeContextCurrent()

	return window, nil
}

func HandleWindow(window *glfw.Window) {
	// Clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// Swap buffers, which swaps the front and back buffers of the window
	// (one used for rendering, the other for drawing)
	window.SwapBuffers()
	// Poll events, which actually allows us to catch events such as the close
	// button, so that the window can close (and, of course, all other events,
	// including input events)
	glfw.PollEvents()
}

func DestroyWindow(window *glfw.Window) {
	window.Destroy()
}
