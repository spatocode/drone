package krypton

type Window interface {
	Title() string
	Resizable(bool)
	Run()
}