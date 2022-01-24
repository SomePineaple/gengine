package ui

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	window  *glfw.Window
	title   string
	width   int
	height  int
	vSync   bool
	resized bool
}

func NewWindow(title string, width int, height int, vSync bool) (w *Window, err error) {
	w = &Window{}

	w.title = title
	w.width = width
	w.height = height
	w.vSync = vSync

	if err = glfw.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize glfw: %v", err)
	}

	w.window, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %v", err)
	}

	if err = gl.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize opengl: %v", err)
	}

	w.window.MakeContextCurrent()

	w.window.SetFramebufferSizeCallback(func(_ *glfw.Window, width int, height int) {
		w.width = width
		w.height = height
		w.resized = true
	})

	return w, nil
}

func (w *Window) Update() {
	glfw.PollEvents()

	if w.resized {
		gl.Viewport(0, 0, int32(w.width), int32(w.height))
		w.resized = false
	}

	w.window.SwapBuffers()
}

func (w *Window) SetClearColor(r float32, g float32, b float32) {
	gl.ClearColor(r, g, b, 1.0)
}

func (w *Window) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w *Window) ShouldClose() bool {
	return w.window.ShouldClose()
}

func (w *Window) Destroy() {
	w.window.Destroy()
}
