package geometry

import (
	"github.com/spatocode/karid"
)

// var _ karid.Geometry = (*triangle)(nil)

type triangle struct {
	size karid.Size
	position karid.Position
	color karid.Color
}

func (tri *triangle) SetSize(size karid.Size) {
	tri.size = size
}

func (tri *triangle) SetPosition(position karid.Position) {
	tri.position = position
}

func (tri *triangle) SetColor(color karid.Color) {
	tri.color = color
}

// NewTriangle creates a new triangle geometry
func NewTriangle() *karid.Geometry {
	return &triangle{}
}
