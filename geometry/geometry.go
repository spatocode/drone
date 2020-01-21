package geometry

import (
	"github.com/spatocode/karid"
)

// Geometry is a parent struct for all geometrical shapes
type Geometry struct {
	size karid.Size
	position karid.Position
	color karid.Color
}

func (geo *geometry) SetSize(size karid.Size) {
	geo.size = size
}

func (geo *geometry) SetPosition(position karid.Position) {
	geo.position = position
}

func (geo *geometry) RotateX(rad float32) {
}

func (geo *geometry) RotateY(rad float32) {
}

func (geo *geometry) RotateZ(rad float32) {
}

func (geo *geometry) TranslateX(rad float32) {
}

func (geo *geometry) TranslateY(rad float32) {
}

func (geo *geometry) TranslateZ(rad float32) {
}

func (geo *geometry) Translate(x, y, z float32) {
}

func (geo *geometry) Rotate(x, y, z float32) {
}

func (geo *geometry) Position() karid.Position {
	return geo.position
}

func (geo *geometry) SetColor(color karid.Color) {
	geo.color = color
}

// New creates a new geometry
func New() karid.Geometry {
	return &geometry{}
}
