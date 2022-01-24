package items

import (
	"github.com/SomePineaple/gengine/engine/shading"
	"github.com/go-gl/mathgl/mgl32"
)

type GameItem struct {
	position mgl32.Vec3
	rotation mgl32.Vec3
	scale    float32
	mesh     *Mesh
}

// NewGameItem Creates a new game item
func NewGameItem(mesh *Mesh) *GameItem {
	return &GameItem{
		mesh:     mesh,
		position: mgl32.Vec3{0, 0, 0},
		rotation: mgl32.Vec3{0, 0, 0},
		scale:    1,
	}
}

// Render Draws the game item to the display
func (gi *GameItem) Render(shaderProgram *shading.ShaderProgram) {
	// TODO: Render game items
}

// SetPosition Sets the position of the game item
func (gi *GameItem) SetPosition(x, y, z float32) {
	gi.position[0] = x
	gi.position[1] = y
	gi.position[2] = z
}

// SetRotation Sets the rotation of this game item
func (gi *GameItem) SetRotation(x, y, z float32) {
	gi.rotation[0] = x
	gi.rotation[1] = y
	gi.rotation[2] = z
}

// Destroy Frees the memory used by this game item
func (gi *GameItem) Destroy() {
	gi.mesh.Destroy()
}
