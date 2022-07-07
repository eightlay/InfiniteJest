package ui

import (
	uidriver "github.com/eightlay/InfiniteJest/iternal/ui"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Window struct
type Window struct {
	window *glfw.Window
}

// Create a window with the given parameters
func CreateWindow(title string, width, height int, resizable bool) (*Window, error) {
	window, err := uidriver.CreateWindow(title, width, height, resizable)
	if err != nil {
		return nil, err
	}
	return &Window{window}, nil
}

// Check if the window is running
func (w *Window) IsRunning() bool {
	return !w.window.ShouldClose()
}

// Clear window
func (*Window) Clear() {
	uidriver.Clear()
}

// Swap buffers and poll events of the window
func (w *Window) Handle() {
	uidriver.HandleWindow(w.window)
}

// Destroy window
func (w *Window) Destroy() {
	uidriver.DestroyWindow(w.window)
}

// Check if key is pressed
func (w *Window) KeyIsPressed(key glfw.Key) bool {
	return uidriver.KeyIsPressed(w.window, key)
}

// Check if mouse is pressed
func (w *Window) MouseIsPressed(button glfw.MouseButton) bool {
	return uidriver.MouseIsPressed(w.window, button)
}

// Set cursor pos callback
func (w *Window) SetCursorPosCallback(cbfun glfw.CursorPosCallback) {
	w.window.SetCursorPosCallback(cbfun)
}

// Set scroll callback
func (w *Window) SetScrollCallback(cbfun glfw.ScrollCallback) {
	w.window.SetScrollCallback(cbfun)
}
