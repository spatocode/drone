package geometry

// Torus represents a torus geometry
type Torus struct {
	Geometry
	radius float32
	tubeRadius float32
}

// SetRadius sets the radius of this geometry to rad
func (t *Torus) SetRadius(rad float32) {
	t.radius = rad
}

// SetTubeRadius sets the tube radius of this geometry to tub
func (t *Torus) SetTubeRadius(tub float32) {
	t.tubeRadius = tub
}

// NewTorus creates a new torus geometry
// with a radius rad and tube radius tub
func NewTorus(rad float32, tub float32) *Torus {
	return &Torus{
		radius: rad,
		tubeRadius: tub,
	}
}