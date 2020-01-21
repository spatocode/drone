package geometry

import (
	"github.com/spatocode/karid"
)

type Triangle struct {
	geometry
}

// NewTriangle creates a new triangle geometry
func NewTriangle() karid.Geometry {
	return &triangle{}
}
