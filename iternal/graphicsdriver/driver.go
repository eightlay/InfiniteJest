package graphicsdriver

import "fmt"

var initialized bool = false

// Initialize graphics driver
func Init() error {
	if !IsInitialized() {
		InitGL()
		InitGLFW()
		initialized = true
		return nil
	}
	return fmt.Errorf("graphics are already initialized")
}

// Terminate graphics driver
func Terminate() error {
	if IsInitialized() {
		TerminateGLFW()
		initialized = false
		return nil
	}
	return fmt.Errorf("graphics are already terminated")
}

// Check if graphics driver is initialized
func IsInitialized() bool {
	return initialized
}
