package app

import (
	"github.com/spatocode/krypton"
)

// App represents a krypton app
type App struct {
	window	krypton.Window
}

func (app *App) Run() {
	app.window.Run()
}

func (app *App) RegisterWindow(obj krypton.Window) {
	app.window = obj
}

// Init initializes a krypton app
func Init() *App {
	app := new(App)
	return app
}
