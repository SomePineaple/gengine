package example_game

import (
	"github.com/SomePineaple/gengine/engine/items"
	"github.com/SomePineaple/gengine/engine/shading"
	"github.com/SomePineaple/gengine/ui"
	"github.com/SomePineaple/gengine/utils"
	"log"
	"os"
)

var window *ui.Window

const (
	WindowWidth  = 800
	WindowHeight = 600
	VSync        = true
)

var triangle *items.Mesh
var shaderProgram *shading.ShaderProgram

func StartGame() {
	var err error
	window, err = ui.NewWindow("Example Game", WindowWidth, WindowHeight, VSync)
	if err != nil {
		log.Fatalln("Failed to create window:", err)
	}
	defer window.Destroy()

	window.SetClearColor(0.4, 0.3, 0.7)

	triangle = items.NewMesh([]float32{
		0.0, 0.5, 0.0,
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
	}, []float32{
		0, 0, 0,
	}, []float32{
		0, 0, 0, 0,
	}, []int32{
		0, 1, 2,
	}, "triangle")
	defer triangle.Destroy()

	vertexShader, err := os.ReadFile("data/example_game/shaders/basic.vsh")
	if err != nil {
		log.Fatalln("Could not get vertex shader code:", err)
	}

	fragmentShader, err := os.ReadFile("data/example_game/shaders/basic.fsh")
	if err != nil {
		log.Fatalln("Could not get fragment shader code:", err)
	}

	shaderProgram, err = shading.NewShaderProgram(string(vertexShader), string(fragmentShader))
	if err != nil {
		log.Fatalln("Could not create shader program:", err)
	}
	defer shaderProgram.Destroy()

	gameLoop()
}

func gameLoop() {
	for !window.ShouldClose() {
		update()
		render()
	}
}

func update() {
	utils.PrintErrors()
}

func render() {
	window.Clear()

	shaderProgram.Bind()

	triangle.Render()

	shaderProgram.Unbind()

	window.Update()
}
