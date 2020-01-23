package karid

type Scene interface {
	Object() []Object
	Background() interface{}
}
