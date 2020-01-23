package geometry

// Circle represents a circle geometry
type Circle struct {
	Geometry
	radius float32
	segments int32
}

// SetRadius sets the radius of this geometry to rad
func (c *Circle) SetRadius(rad float32) {
	c.radius = rad
}

// SetSegment sets the number segments of this geometry
func (c *Circle) SetSegment(seg int32) {
	c.segments = seg
}

// NewCircle creates a new circle geometry
// with a radius rad and seg number of segments
func NewCircle(rad float32, seg int32) *Circle {
	return &Circle{radius: rad, segments: seg}
}
