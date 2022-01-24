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

// NewWindow Creates a new window
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

	if vSync {
		glfw.SwapInterval(1)
	}

	w.window.SetFramebufferSizeCallback(func(_ *glfw.Window, width int, height int) {
		w.width = width
		w.height = height
		w.resized = true
	})

	return w, nil
}

// Update Polls events, checks if the window has been resized, and swaps window buffers
func (w *Window) Update() {
	glfw.PollEvents()

	if w.resized {
		gl.Viewport(0, 0, int32(w.width), int32(w.height))
		w.resized = false
	}

	w.window.SwapBuffers()
}

// SetClearColor Sets the opengl clear color, this is the color shown behind all the stuff being rendered
func (w *Window) SetClearColor(r float32, g float32, b float32) {
	gl.ClearColor(r, g, b, 1.0)
}

// Clear Clears the opengl color buffer and the depth buffer
func (w *Window) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// ShouldClose Returns weather the window should close, returns true when the user clicks the x button
func (w *Window) ShouldClose() bool {
	return w.window.ShouldClose()
}

// Show Shows the window on the screen
func (w *Window) Show() {
	w.window.Show()
}

func (w *Window) Hide() {
	w.window.Hide()
}

func (w *Window) Destroy() {
	w.window.Destroy()
}
