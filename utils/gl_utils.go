package utils

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"log"
	"strings"
)

func PrintErrors() {
	for err := gl.GetError(); err != gl.NO_ERROR; err = gl.GetError() {
		log.Println("OpenGL Error:", err)
	}
}

func CreateShader(source string, shaderType uint32) (uint32, error) {
	shaderID := gl.CreateShader(shaderType)
	if shaderID == 0 {
		return shaderID, fmt.Errorf("failed to create shader of type %v", shaderType)
	}

	csources, free := gl.Strs(source)
	gl.ShaderSource(shaderID, 1, csources, nil)
	free()
	gl.CompileShader(shaderID)

	var status int32
	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logLength)

		compileLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderID, logLength, nil, gl.Str(compileLog))

		return 0, fmt.Errorf("failed to compile %v: %v", source, compileLog)
	}

	return shaderID, nil
}
