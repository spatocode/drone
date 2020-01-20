package scene

import (
	"github.com/spatocode/karid"
)

var _ karid.Scene = (*scene)(nil)

type scene struct {
	scale float32
	camera karid.Camera
	// light  karid.Light
	// material karid.Material
	// mesh karid.Mesh
	geometry karid.Geometry
}

func (sc *scene) SetContent(content karid.Geometry) {
	sc.geometry = content
}

func (sc *scene) SetCamera(cam karid.Camera) {
	sc.camera = cam
}

func (sc *scene) Content() karid.Geometry {
	return sc.geometry
}

func New() karid.Scene {
	sc := &scene{scale: 1.0}
	return sc
}
