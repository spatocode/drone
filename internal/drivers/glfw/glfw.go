package glfw

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/spatocode/drone"
	_ "github.com/spatocode/drone/internal/drivers"
)

type Driver struct {
	window *Window
}

// Ensures we execute on current OS thread
func init() {
	runtime.LockOSThread()
}

// CreateWindow creates a new window using GLFW
func (driver *Driver) CreateWindow(title string, width, height int) *Window {
	name := "Drone"
	if title != "" {
		name = title
	}

	win, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		drone.LogError("Failed to create a window", err)
	}
	
	window := &Window{
		title: name,
		width: width,
		height: height,
	}

	driver.window = window

	win.MakeContextCurrent()
	defer glfw.Terminate()

	for !win.ShouldClose() {
		win.SwapBuffers()
   		glfw.PollEvents()
	}

	return driver.window
}

// StartDriver initialize and start the underlying driver
func StartDriver() *Driver {
	driver := new(Driver)
	err := glfw.Init()
	if err != nil {
		drone.LogError("Failed to initialize GLFW", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	return driver
}
