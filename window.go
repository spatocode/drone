package karid

type Window interface {
	Title() string
	Resizable(bool)
	Run()
}