package app

import (
	"github.com/spatocode/drone/internal/drivers/glfw"
)

// App represents a drone app
type App struct {
	driver *glfw.Driver
}

// CreateWindow creates a new window
func (app *App) CreateWindow(title string, width int, height int) {
	app.driver.CreateWindow(title, width, height)
}

// Launch starts a Drone app
func Launch() *App {
	driver := glfw.StartDriver()
	app := &App{driver: driver}
	return app
}
