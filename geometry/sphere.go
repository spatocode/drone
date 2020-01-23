package geometry

// Sphere represents a sphere geometry
type Sphere struct {
	Geometry
	radius float32
}

// SetRadius sets the radius of this geometry
func (s *Sphere) SetRadius(rad float32) {
	s.radius = rad
}

// NewSphere creates a new sphere geometry with radius rad
func NewSphere(rad float32) *Sphere {
	return &Sphere{radius: rad}
}
