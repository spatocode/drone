package window

import (
	"runtime"

	"github.com/spatocode/krypton"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var _ krypton.Window = (*window)(nil)

type window struct {
	width, height	int
	resizable	bool
	title		string
	fullscreen	bool
	view	*glfw.Window
	visible	bool
}

func init() {
	// Make sure OpenGL function calls runs in the same thread.
	runtime.LockOSThread()
}

func (win *window) Title() string{
	return win.title
}

func (win *window) SetTitle(title string) {
	win.title = title
	win.view.SetTitle(title)
}

func (win *window) Resizable(resize bool)  {
	if resize {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	win.resizable = resize
}

func (win *window) FullScreen(bo bool) {

}

func (win *window) destroy() {
}

func (win *window) position() {	
}

func (win *window) Hide() {
	win.visible = false
	win.view.Hide()
}

func (win *window) Show() {
	win.visible = true
	win.view.Show()
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
	win.Show()
	view := win.view
	for !view.ShouldClose() {
		view.SwapBuffers()
		glfw.PollEvents()
	}

	defer glfw.Terminate()
}

// New creates a new window
func New(title string, width, height int) krypton.Window {
	err := glfw.Init()
	if err != nil {
		krypton.LogError("Failed to initialize GLFW", err)
	}

	glfw.WindowHint(glfw.Visible, glfw.False)

	name := "Krypton"
	if title != "" {
		name = title
	}

	win, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		krypton.LogError("Failed to create a window", err)
	}

	window := &window{
		view:	win,
		title:  name,
		width:  width,
		height: height,
	}

	win.MakeContextCurrent()

	return window
}
