package karid

// Geometry defines the behaviour of a geometrical object
type Geometry interface {
	SetColor(Color)
	SetSize(Size)
	SetPosition(Position)
}

// Color represents the RGBA color of an object
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// Position represents the x and y coordinates of an object
type Position struct {
	X int
	Y int
}

// Size represents the width and heigth of an object
type Size struct {
	Width int
	Height int
}

// NewColor initialize a new color of an object
func NewColor(r, g, b, a uint8) Color {
	return Color{r, g, b, a}
}

// NewPosition initialize a new position of an object
func NewPosition(x, y int) Position {
	return Position{x, y}
}

// NewSize initialize a new size of an object
func NewSize(width, height int) Size {
	return Size{width, height}
}