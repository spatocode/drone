package glfw

import (
	"github.com/spatocode/drone"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var _ drone.Window = (*Window)(nil)

type Window struct {
	width     int
	height    int
	resizable bool
	title     string
}

func (window *Window) Title() string{
	return window.title
}

func (window *Window) Resizable(bo bool) {
	window.resizable = bo
}

func (window *Window) Run() {
	if window.resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
}
