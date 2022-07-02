package tests

import (
	"runtime"
	"testing"

	"github.com/eightlay/InfiniteJest/ui"
)

func TestWindow(t *testing.T) {
	runtime.LockOSThread()

	window := ui.CreateWindow("Test", 800, 600, true)

	window.IsRunning()
	window.Handle()

}
