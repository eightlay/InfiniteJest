package graphicsdriver

import "fmt"

// Indicate if graphics driver was initialized
var initialized bool = false
var engine3D bool = false

// Initialize graphics driver
func Init(use3D bool) error {
	err := error(nil)

	if !IsInitialized() {
		err = InitGL()
		if err != nil {
			return err
		}

		err = InitGLFW()
		if err != nil {
			return err
		}

		initialized = true
		engine3D = use3D
	} else {
		err = fmt.Errorf("graphics driver is already initialized")
	}

	return err
}

// Terminate graphics driver
func Terminate() error {
	if IsInitialized() {
		TerminateGLFW()
		initialized = false
		return nil
	}
	return fmt.Errorf("graphics driver is not initialized")
}

// Check if graphics driver is initialized
func IsInitialized() bool {
	return initialized
}

// Check if graphics driver is for 3D
func Is3D() bool {
	return engine3D
}
