package karid

type Scene interface {
	SetContent(Geometry)
	Content() Geometry
}
