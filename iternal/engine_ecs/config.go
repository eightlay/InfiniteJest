package engine_ecs

import (
	"github.com/eightlay/InfiniteJest/iternal/graphics/shaders"
	"github.com/eightlay/InfiniteJest/iternal/ui"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type ConfigEngineECS struct {
	shader   *shaders.Shader
	camera   *ui.Camera
	window_w int
	window_h int
	window   *glfw.Window
}

func CreateConfigEngineECS(
	shader *shaders.Shader, camera *ui.Camera, window_w, window_h int, window *glfw.Window,
) (*ConfigEngineECS, error) {
	return &ConfigEngineECS{shader, camera, window_w, window_h, window}, nil
}

func (c *ConfigEngineECS) Shader() *shaders.Shader {
	return c.shader
}

func (c *ConfigEngineECS) Camera() *ui.Camera {
	return c.camera
}

func (c *ConfigEngineECS) WindowW() int {
	return c.window_w
}

func (c *ConfigEngineECS) WindowH() int {
	return c.window_h
}

func (c *ConfigEngineECS) Window() *glfw.Window {
	return c.window
}
