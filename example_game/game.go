package example_game

import (
	"github.com/SomePineaple/gengine/ui"
	"log"
)

var window *ui.Window

const (
	WindowWidth  = 800
	WindowHeight = 600
	VSync        = true
)

func StartGame() {
	var err error
	window, err = ui.NewWindow("Example Game", WindowWidth, WindowHeight, VSync)
	if err != nil {
		log.Fatalln("Failed to create window:", err)
	}
	defer window.Destroy()

	window.SetClearColor(0.4, 0.3, 0.7)

	gameLoop()
}

func gameLoop() {
	for !window.ShouldClose() {
		update()
		render()
	}
}

func update() {

}

func render() {
	window.Clear()
	window.Update()
}
