package geometry

import (
	"github.com/spatocode/karid"
)

type Circle struct {
	geometry
}

// NewCircle creates a new circle geometry
func NewCircle() karid.Geometry {
	return &circle{}
}
