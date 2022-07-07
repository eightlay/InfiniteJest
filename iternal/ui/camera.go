package ui

import (
	"math"

	glm "github.com/go-gl/mathgl/mgl32"
)

var directions = [4]string{"forward", "backward", "left", "right"}

// Default camera values
const (
	YAW         = -90.0
	SPEED       = 2.5
	SENSITIVITY = 0.1
	ZOOM        = 45.0
)

type Camera struct {
	// Attributes
	position glm.Vec3
	front    glm.Vec3
	up       glm.Vec3
	right    glm.Vec3
	worldUp  glm.Vec3

	// Euler angles
	yaw   float32
	pitch float32

	// Options
	movementSpeed    float32
	mouseSensitivity float32
	zoom             float32
}

func CreateCamera(position, up [3]float32, yaw, pitch float32) (*Camera, error) {
	camera := &Camera{
		position: position,
		front:    glm.Vec3{0, 0, -1},
		worldUp:  up,

		yaw:   yaw,
		pitch: pitch,

		movementSpeed:    SPEED,
		mouseSensitivity: SENSITIVITY,
		zoom:             ZOOM,
	}
	camera.updateCameraVectors()
	return camera, nil
}

func (c *Camera) ViewMatrix() glm.Mat4 {
	return glm.LookAtV(c.position, c.position.Add(c.front), c.up)
}

func (c *Camera) Zoom() float32 {
	return glm.DegToRad(c.zoom)
}

func (c *Camera) MoveKeyboard(direction int, dt float32) {
	velocity := c.movementSpeed * dt
	switch directions[direction] {
	case "forward":
		c.position = c.position.Add(c.front.Mul(velocity))
	case "backward":
		c.position = c.position.Sub(c.front.Mul(velocity))
	case "left":
		c.position = c.position.Sub(c.right.Mul(velocity))
	case "right":
		c.position = c.position.Add(c.right.Mul(velocity))
	}
}

func (c *Camera) MoveMouse(xoffset, yoffset float32) {
	xoffset *= c.mouseSensitivity
	yoffset *= c.mouseSensitivity

	c.yaw += xoffset
	c.pitch += yoffset

	if c.pitch > 89 {
		c.pitch = 89
	} else if c.pitch < -89 {
		c.pitch = -89
	}

	c.updateCameraVectors()
}

func (c *Camera) MoveScroll(yoffset float32) {
	c.zoom -= yoffset

	if c.zoom < 1 {
		c.zoom = 1
	} else if c.zoom > 45 {
		c.zoom = 45
	}

	c.updateCameraVectors()
}

func (c *Camera) updateCameraVectors() {
	yawRad := float64(glm.DegToRad(c.yaw))
	pitchRad := float64(glm.DegToRad(c.pitch))
	cosPitch := math.Cos(pitchRad)

	front := glm.Vec3{
		float32(math.Cos(yawRad * cosPitch)),
		float32(math.Sin(pitchRad)),
		float32(math.Sin(yawRad * cosPitch)),
	}
	c.front = front.Normalize()

	c.right = c.front.Cross(c.worldUp).Normalize()
	c.up = c.right.Cross(c.front).Normalize()
}
