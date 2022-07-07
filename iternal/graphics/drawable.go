package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Drawable object
type Drawable struct {
	id             uint32
	verticesNumber int32
}

// Draw object to the window
func (d *Drawable) Draw() {
	gl.BindVertexArray(d.id)
	gl.DrawElementsWithOffset(gl.TRIANGLES, int32(d.verticesNumber), gl.UNSIGNED_INT, 0)
}

// Create new drawable with the given vertices and its indices
func CreateDrawable(verts []float32, indices []uint32) (*Drawable, error) {
	var vbo, vao, ebo uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(verts), gl.Ptr(verts), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	// Position attribute
	// 5th argument = (3 + 2) * sizeof(float32)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 20, 0)
	gl.EnableVertexAttribArray(0)

	// Texture coord attribute
	// 6th argument = 3 * sizeof(float32)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 20, 12)
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return &Drawable{vao, int32(len(indices))}, nil
}
