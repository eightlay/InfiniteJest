package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Drawable struct {
	vao            uint32
	verticesNumber int32
}

func (d *Drawable) Draw() {
	gl.BindVertexArray(d.vao)
	gl.DrawElementsWithOffset(gl.TRIANGLES, int32(d.verticesNumber), gl.UNSIGNED_INT, 0)
}

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

	// 5th argument = 3 * sizeof(float32)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 12, 0)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return &Drawable{vao, int32(len(indices))}, nil
}
