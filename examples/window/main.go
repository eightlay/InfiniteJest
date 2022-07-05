package main

import (
	"fmt"
	"runtime"

	"github.com/eightlay/InfiniteJest/ui"
)

func main() {
	// Lock thread
	runtime.LockOSThread()

	// Init graphics driver
	err := ui.Init()
	if err != nil {
		panic(fmt.Errorf("could not initilize graphics driver: %v", err))
	}
	defer ui.Terminate()

	// Create window
	window, err := ui.CreateWindow("Test", 800, 600, true)
	if err != nil {
		panic(fmt.Errorf("could not create window: %v", err))
	}
	defer window.Destroy()

	// Handle window
	for window.IsRunning() {
		window.Handle()
	}
}
