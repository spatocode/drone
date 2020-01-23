package geometry

// Cylinder represents a cylinder geometry
type Cylinder struct {
	Geometry
	height        float32
	topRadius     float32
	bottomRadius  float32
	radialSegment int
	heightSegment int
	open bool
}


// SetHeight sets the height of this geometry
func (c *Cylinder) SetHeight(height float32) {
	c.height = height
}

// SetTopRadius sets the top radius of this geometry
func (c *Cylinder) SetTopRadius(rad float32) {
	c.topRadius = rad
}

// SetBottomRadius sets the bottom radius of this geometry
func (c *Cylinder) SetBottomRadius(rad float32) {
	c.bottomRadius = rad
}

// Open sets the ends of this geometry to open or close
func (c *Cylinder) Open(open bool) {
	c.open = open
}

// IsOpen indicates whether the ends of this geometry is open or closed
func (c *Cylinder) IsOpen() bool {
	return c.open
}

// NewCylinder creates a new cylinder geometry
// with height, topRadius, bottomRadius and open
func NewCylinder(height, topRad, bottRad float32, open bool) *Cylinder {
	return &Cylinder{
		height: height, 
		topRadius: topRad, 
		bottomRadius: bottRad,
		open: open,
	}
}
