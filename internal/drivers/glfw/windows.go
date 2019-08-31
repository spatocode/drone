package glfw

import (
	_"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	width     int
	height    int
	resizable bool
	title     string
}
