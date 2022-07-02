package main

import (
	"runtime"

	"github.com/eightlay/InfiniteJest/ui"
)

func main() {
	runtime.LockOSThread()

	window := ui.CreateWindow("Test", 800, 600, true)

	for window.IsRunning() {
		window.Handle()
	}
}
