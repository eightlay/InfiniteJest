package ui

import "github.com/eightlay/InfiniteJest/iternal/graphicsdriver"

// Initialize graphics driver
func Init() error {
	return graphicsdriver.Init()
}

// Terminate graphics driver
func Terminate() {
	graphicsdriver.Terminate()
}
