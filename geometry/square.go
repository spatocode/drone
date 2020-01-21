package geometry

import (
	"github.com/spatocode/karid"
)

type Square struct {
	geometry
}

// NewSquare creates a new square geometry
func NewSquare() karid.Geometry {
	return &square{}
}
