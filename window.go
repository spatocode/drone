package karid

type Window interface {
	Title() string
	Resizable(bool)
	SetRenderer(Renderer)
	Show()
	Run()
}