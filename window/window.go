package window

import (
	"github.com/spatocode/krypton"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var _ krypton.Window = (*window)(nil)

type window struct {
	width, height	int
	resizable	bool
	title		string
	fullscreen	bool
	view	glfw.Window
}

func (win *window) Title() string{
	return win.title
}

func (win *window) Resizable(bo bool)  {
	switch bo {
	case true:
		glfw.windowHint(glfw.Resizable, glfw.True)
	default:
		glfw.windowHint(glfw.Resizable, glfw.False)
	}
	win.resizable = bo
}

func (win *window) FullScreen(bo bool) {
	switch bo {
	case true:
		glfw.windowHint(glfw.Resizable, glfw.True)
	default:
		glfw.windowHint(glfw.Resizable, glfw.False)
	}
	win.fullscreen = bo
}

func (win *window) Run() {
}

func (win *window) destroy() {
}

func (win *window) position() {	
}

func (win *window) Hide() {	
}

func (win *window) Show() {	
}

func (win *window) Restore() {	
}

func (win *window) SetPosition() {	
}

func (win *window) SetBrightness() {	
}

func (win *window) SetIcon() {
	
}

func (win *window) Run() {
	for !win.ShouldClose() {
		win.SwapBuffers()
		glfw.PollEvents()
	}
}

// New creates a new window
func New(title string, width, height int) krypton.Window {
	name := "Krypton"
	if title != "" {
		name = title
	}

	win, err := glfw.Createwindow(width, height, name, nil, nil)
	if err != nil {
		krypton.LogError("Failed to create a window", err)
		return
	}

	window := &window{
		view:	win
		title:  name,
		width:  width,
		height: height,
	}

	win.MakeContextCurrent()

	return window
}
