package karid

type Scene interface {
	Object() []Object
	AddObject(Object)
	Background() Colorable
	SetBackground(Colorable)
}
