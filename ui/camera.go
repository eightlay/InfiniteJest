package ui

import (
	uidriver "github.com/eightlay/InfiniteJest/iternal/ui"
	glm "github.com/go-gl/mathgl/mgl32"
)

type CameraMovement int

// Possible camera movements
const (
	CameraForward CameraMovement = iota
	CameraBackward
	CameraLeft
	CameraRight
)

type Camera struct {
	camera *uidriver.Camera
}

func CreateCamera(position, up [3]float32, yaw, pitch float32) (*Camera, error) {
	camera, err := uidriver.CreateCamera(position, up, yaw, pitch)
	if err != nil {
		return nil, err
	}
	return &Camera{camera}, nil
}

func (c Camera) ViewMatrix() glm.Mat4 {
	return c.camera.ViewMatrix()
}

func (c Camera) Zoom() float32 {
	return c.camera.Zoom()
}

func (c *Camera) MoveKeyboard(direction CameraMovement, dt float32) {
	c.camera.MoveKeyboard(int(direction), dt)
}

func (c *Camera) MoveMouse(xoffset, yoffset float32) {
	c.camera.MoveMouse(xoffset, yoffset)
}

func (c *Camera) MoveScroll(yoffset float32) {
	c.camera.MoveScroll(yoffset)
}
