package krypton

import (
	"github.com/spatocode/krypton"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// App represents a krypton app
type App struct {
	window	krypton.Window
}

func (app *App) Run() {
}

// Init initializes a krypton app
func Init() *App {
	app := new(App)
	err := glfw.Init()
	if err != nil {
		krypton.LogError("Failed to initialize GLFW", err)
		return
	}
	return app
}
