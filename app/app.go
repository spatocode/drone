package app

import (
	"github.com/spatocode/karid"
)

// App represents a karid app
type App struct {
	window	karid.Window
}

func (app *App) Run() {
	app.window.Run()
}

func (app *App) RegisterWindow(obj karid.Window) {
	app.window = obj
}

// Init initializes a karid app
func Init() *App {
	app := new(App)
	return app
}
