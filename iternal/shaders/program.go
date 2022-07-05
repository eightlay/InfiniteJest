package shaders

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type GLProgram struct {
	program uint32
}

func (p *GLProgram) Use() {
	gl.UseProgram(p.program)
}

func GetBasicProgram() (*GLProgram, error) {
	return createProgram(VertexShaderSource, FragmentShaderSource)
}

func createProgram(vertexShaderSource, fragmentShaderSource string) (*GLProgram, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return &GLProgram{0}, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return &GLProgram{0}, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return &GLProgram{0}, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &GLProgram{program}, nil
}
