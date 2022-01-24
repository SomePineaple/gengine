package shading

import (
	"fmt"
	"github.com/SomePineaple/gengine/utils"
	"github.com/go-gl/gl/all-core/gl"
	"strings"
)

type ShaderProgram struct {
	fragmentShaderCode string
	vertexShaderCode   string

	programID        uint32
	fragmentShaderID uint32
	vertexShaderID   uint32

	uniforms map[string]int
}

// NewShaderProgram Creates a new shader program with the given vertex and fragment shaders
func NewShaderProgram(vertexShaderCode string, fragmentShaderCode string) (sh *ShaderProgram, err error) {
	sh = &ShaderProgram{}

	sh.vertexShaderCode = vertexShaderCode
	sh.fragmentShaderCode = fragmentShaderCode

	sh.programID = gl.CreateProgram()
	if sh.programID == 0 {
		return nil, fmt.Errorf("failed to create shader program")
	}

	sh.vertexShaderID, err = utils.CreateShader(vertexShaderCode, gl.VERTEX_SHADER)
	if err != nil {
		return nil, fmt.Errorf("failed to create vertex shader: %v", err)
	}

	sh.fragmentShaderID, err = utils.CreateShader(fragmentShaderCode, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, fmt.Errorf("failed to create fragment shader: %v", err)
	}

	err = sh.link()
	if err != nil {
		return nil, fmt.Errorf("failed to link shader program: %v", err)
	}

	return sh, nil
}

// Bind Tells opengl to use this shader program
func (sh *ShaderProgram) Bind() {
	gl.UseProgram(sh.programID)
}

// Unbind Tells opengl not to use this shader program
func (sh *ShaderProgram) Unbind() {
	gl.UseProgram(0)
}

// Destroy Cleanup the program, free memory and such
func (sh *ShaderProgram) Destroy() {
	sh.Unbind()
	if sh.programID != 0 {
		gl.DeleteProgram(sh.programID)
	}
}

func (sh *ShaderProgram) link() error {
	gl.LinkProgram(sh.programID)

	var status int32
	gl.GetProgramiv(sh.programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(sh.programID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(sh.programID, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(sh.vertexShaderID)
	gl.DeleteShader(sh.fragmentShaderID)

	return nil
}
