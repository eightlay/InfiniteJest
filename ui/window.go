package ui

import (
	"github.com/eightlay/InfiniteJest/iternal/uidriver"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Window struct
type Window struct {
	window *glfw.Window
}

func (w *Window) IsRunning() bool {
	return !w.window.ShouldClose()
}

func (*Window) Clear() {
	uidriver.Clear()
}

func (w *Window) Handle() {
	uidriver.HandleWindow(w.window)
}

func (w *Window) Destroy() {
	uidriver.DestroyWindow(w.window)
}

func CreateWindow(title string, width, height int, resizable bool) (*Window, error) {
	window, err := uidriver.CreateWindow(title, width, height, resizable)
	if err != nil {
		return nil, err
	}
	return &Window{window}, nil
}
