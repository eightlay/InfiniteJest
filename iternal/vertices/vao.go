package vertices

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VAO struct {
	drawable uint32
}

func (vao *VAO) Draw() {
	gl.BindVertexArray(vao.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func CreateVAO(verts []float32) (*VAO, error) {
	var vbo, vao uint32

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(verts), gl.Ptr(verts), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// 5th argument = 3 * sizeof(float32)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 12, 0)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return &VAO{vao}, nil
}
