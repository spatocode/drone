package geometry

import (
	"github.com/spatocode/karid"
)

// Geometry is a parent struct for all geometrical shapes
type Geometry struct {
	position karid.Position
	color karid.Color
}

func (geo *Geometry) AddObject(obj karid.Object) {
}

// SetPosition sets the position of this geometry
func (geo *Geometry) SetPosition(position karid.Position) {
	geo.position = position
}

// RotateX rotates this geometry along it's x axis
func (geo *Geometry) RotateX(rad float32) {
}

// RotateY rotates this geometry along it's y axis
func (geo *Geometry) RotateY(rad float32) {
}

// RotateZ rotates this geometry along z axis
func (geo *Geometry) RotateZ(rad float32) {
}

// TranslateX translates this geometry along x axis
func (geo *Geometry) TranslateX(rad float32) {
}

// TranslateY translates this geometry along y axis
func (geo *Geometry) TranslateY(rad float32) {
}

// TranslateZ translates this geometry along z axis
func (geo *Geometry) TranslateZ(rad float32) {
}

// Translate translates this geometry along it's x, y, z axis
func (geo *Geometry) Translate(x, y, z float32) {
}

// Rotate rotates this geometry along it's x, y, z axis
func (geo *Geometry) Rotate(x, y, z float32) {
}

// Position returns the position of this geometry
func (geo *Geometry) Position() karid.Position {
	return geo.position
}

// SetColor sets the color of this geometry
func (geo *Geometry) SetColor(color karid.Color) {
	geo.color = color
}

// New creates a new geometry
func New() *Geometry {
	return &Geometry{}
}
