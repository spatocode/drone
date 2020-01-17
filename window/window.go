package window

import (
	"runtime"

	"github.com/spatocode/karid"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var _ karid.Window = (*window)(nil)

type window struct {
	width, height	int
	resizable	bool
	title		string
	fullscreen	bool
	view	*glfw.Window
	visible	bool
	renderer karid.Renderer
	running bool
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

func (win *window) SetRenderer(renderer karid.Renderer) {
	win.renderer = renderer
}

func (win *window) destroy() {
}

func (win *window) position(w *glfw.Window, xpos, ypos int) {	
}

func (win *window) Hide() {
	win.visible = false
	win.view.Hide()
}

func (win *window) Show() {
	win.visible = true
	win.view.Show()
	if win.renderer && !win.running {
		win.renderer.Run()
		win.running = true
	}
}

func (win *window) Restore() {	
}

func (win *window) SetPosition() {	
}

func (win *window) SetBrightness() {	
}

func (win *window) SetIcon() {
}

func (win *window) close(w *glfw.Window) {
}

func (win *window) charInput(w *glfw.Window, char rune) {
}

func (win *window) charModInput(w *glfw.Window, char rune, mods glfw.ModifierKey) {
}

func (win *window) mouseClick(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
}

func (win *window) resize(w *glfw.Window, width, height int) {
}

func (win *window) scroll(w *glfw.Window, xoff, yoff float64) {
}

func (win *window) keyPress(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
}

func (win *window) focus(w *glfw.Window, focused bool) {
}

func (win *window) cursorMove(w *glfw.Window, xpos, ypos float64) {
}

func (win *window) frameBuffer(w *glfw.Window, width, height int) {
}

func (win *window) cursorEnter(w *glfw.Window, entered bool) {
}

func (win *window) detectJoystick(joy glfw.Joystick, event glfw.PeripheralEvent) {
}

func (win *window) refresh(w *glfw.Window) {
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
func New(title string, width, height int) karid.Window {
	err := glfw.Init()
	if err != nil {
		karid.LogError("Failed to initialize GLFW", err)
	}

	glfw.WindowHint(glfw.Visible, glfw.False)

	name := "karid"
	if title != "" {
		name = title
	}

	win, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		karid.LogError("Failed to create a window", err)
	}

	window := &window{
		view:	win,
		title:  name,
		width:  width,
		height: height,
	}

	win.MakeContextCurrent()
	win.SetCloseCallback(window.close)
	win.SetCharCallback(window.charInput)
	win.SetMouseButtonCallback(window.mouseClick)
	win.SetSizeCallback(window.resize)
	win.SetScrollCallback(window.scroll)
	win.SetRefreshCallback(window.refresh)
	win.SetPosCallback(window.position)
	win.SetKeyCallback(window.keyPress)
	win.SetFocusCallback(window.focus)
	win.SetCursorPosCallback(window.cursorMove)
	win.SetCursorEnterCallback(window.cursorEnter)
	win.SetCharModsCallback(window.charModInput)
	win.SetFramebufferSizeCallback(window.frameBuffer)
	glfw.SetJoystickCallback(window.detectJoystick)

	return window
}
