package geometry

import (
	"github.com/spatocode/karid"
)

type Cylinder struct {
	geometry
}

// NewCylinder creates a new cylinder geometry
func NewCylinder() karid.Geometry {
	return &cylinder{}
}
