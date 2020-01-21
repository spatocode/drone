package scene

import (
	"github.com/spatocode/karid"
)

type scene struct {
	scale float32
	camera karid.Camera
	light  karid.Light
	object []karid.Object
	background interface{}
	fog Fog
}

func (sc *scene) AddObject(obj karid.Object) {
	sc.object = append(sc.object, obj)
}

func (sc *scene) SetCamera(cam karid.Camera) {
	sc.camera = cam
}

func (sc *scene) Background() interface{} {
	return sc.background
}

func (sc *scene) SetBackground(bg interface{}) {
	sc.background = bg
}

func (sc *scene) SetFog(fog Fog) {
	sc.fog = fog
}

func (sc *scene) Clear() {
}

func (sc *scene) Object() karid.Object {
	return sc.object
}

// New creates a new scene object
func New() karid.Scene {
	sc := &scene{scale: 1.0}
	return sc
}
