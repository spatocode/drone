package geometry

import (
	"github.com/spatocode/karid"
)

type Cube struct {
	geometry
}

// NewCube creates a new cube geometry
func NewCube() karid.Geometry {
	return &cube{}
}
