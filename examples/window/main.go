package main

import (
	"fmt"
	"runtime"

	"github.com/eightlay/InfiniteJest/ui"
)

func main() {
	runtime.LockOSThread()

	ui.Init()
	defer ui.Terminate()

	window, err := ui.CreateWindow("Test", 800, 600, true)
	if err != nil {
		panic(fmt.Errorf("could not create window: %v", err))
	}
	defer window.Destroy()

	for window.IsRunning() {
		window.Handle()
	}
}
