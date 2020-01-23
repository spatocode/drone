package geometry

// Cube represents a cube geometry
type Cube struct {
	Geometry
	width float32
	height float32
	depth float32
}

// SetDepth sets the width of this geometry
func (c *Cube) SetWidth(wid float32) {
	c.width = wid
}

// SetDepth sets the height of this geometry
func (c *Cube) SetHeight(hei float32) {
	c.height = hei
}

// SetDepth sets the depth of this geometry
func (c *Cube) SetDepth(dep float32) {
	c.depth = dep
}

// NewCube creates a new cube geometry
// with a width wid, height hei and depth dep
func NewCube(wid float32, hei float32, dep float32) *Cube {
	return &Cube{width: wid, height: hei, depth: dep}
}
