package tests

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/eightlay/InfiniteJest/ui"
)

func TestWindow(t *testing.T) {
	runtime.LockOSThread()

	ui.Init()
	defer ui.Terminate()

	window, err := ui.CreateWindow("Test", 800, 600, true)
	if err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %v", err))
	}
	defer window.Destroy()

	window.IsRunning()
	window.Handle()

}
