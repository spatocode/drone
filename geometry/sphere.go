package geometry

import (
	"github.com/spatocode/karid"
)

type Sphere struct {
	geometry
}

// NewSphere creates a new sphere geometry
func NewSphere() karid.Geometry {
	return &sphere{}
}
