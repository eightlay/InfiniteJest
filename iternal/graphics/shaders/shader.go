package shaders

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// OpenGL program to use for rendering
type Shader struct {
	id uint32
}

// Installs a Shader's program object as part of current rendering state
func (s *Shader) Use() {
	gl.UseProgram(s.id)
}

// Specify the value of a boolean uniform variable
func (s *Shader) SetBool(name string, value bool) {
	var v int32 = 0

	if value {
		v = 1
	}

	gl.Uniform1i(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), v)
}

// Specify the value of a integer uniform variable
func (s *Shader) SetInt(name string, value int32) {
	gl.Uniform1i(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), value)
}

// Specify the value of a float uniform variable
func (s *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(s.id, gl.Str(name+"\x00")), value)
}

// Create default shader
func GetDefaultShader() (*Shader, error) {
	return createShader(VertexShaderSource, FragmentShaderSource)
}

// Create new program from the given shaders' source strings
func createShader(vertexShaderSource, fragmentShaderSource string) (*Shader, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return &Shader{0}, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return &Shader{0}, err
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

		return &Shader{0}, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Shader{program}, nil
}

// Compile shader by the given source string and shahder type
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	cstrs, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, cstrs, nil)
	free()

	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
