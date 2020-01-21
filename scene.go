package karid

type Scene interface {
	Object()
	Background() interface{}
}
