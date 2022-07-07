package ui

import "github.com/eightlay/InfiniteJest/iternal/graphicsdriver"

// Initialize graphics driver
func Init(use3D bool) error {
	return graphicsdriver.Init(use3D)
}

// Terminate graphics driver
func Terminate() {
	graphicsdriver.Terminate()
}
