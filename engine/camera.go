package engine

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Camera struct {
	position mgl32.Vec3
	rotation mgl32.Vec3
}

// NewCamera Creates a new camera object
func NewCamera() *Camera {
	return &Camera{
		position: mgl32.Vec3{0, 0, 0},
		rotation: mgl32.Vec3{0, 0, 0},
	}
}

// SetPosition Sets the camera position
func (c *Camera) SetPosition(x, y, z float32) {
	c.position[0] = x
	c.position[1] = y
	c.position[2] = z
}

// MovePosition Moves the camera, X is forward/backward, Y is up/down, and Z is side to side.
func (c *Camera) MovePosition(offsetX, offsetY, offsetZ float32) {
	if offsetZ != 0 {
		c.position[0] += float32(math.Sin(float64(mgl32.DegToRad(c.rotation.Y())))) * -1.0 * offsetZ
		c.position[2] += float32(math.Cos(float64(mgl32.DegToRad(c.rotation.Y())))) * offsetZ
	}

	if offsetX != 0 {
		c.position[0] += float32(math.Sin(float64(mgl32.DegToRad(c.rotation.Y()-90)))) * -1.0 * offsetX
		c.position[2] += float32(math.Cos(float64(mgl32.DegToRad(c.rotation.Y()-90)))) * offsetX
	}

	c.position[1] += offsetY
}

// SetRotation Sets the camera rotation
func (c *Camera) SetRotation(x, y, z float32) {
	c.rotation[0] = x
	c.rotation[1] = y
	c.rotation[2] = z
}

// MoveRotation Rotates the camera
func (c *Camera) MoveRotation(offsetX, offsetY, offsetZ float32) {
	c.rotation[0] += offsetX
	c.rotation[1] += offsetY
	c.rotation[2] += offsetZ

	c.rotation[0] = float32(math.Min(math.Max(float64(c.rotation[0]), -90), 90))
}
