package drone

type Window interface {
	Title() string
	Resizable(bool)
	Run()
}