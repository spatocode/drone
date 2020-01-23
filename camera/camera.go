package camera

import (
	"github.com/spatocode/karid"
)

type CameraType int

const (
	Perspective CameraType =  iota
	Orthographi
)

// Camera represents the projection mode of an object
type Camera struct {
	id	   CameraType
	fov	   float32
	aspect float32
	near   float32
	far    float32
	zoom   int
}

// Zoom returns the camera zoom factor
func (cam *Camera) Zoom() int {
	return cam.zoom
}

// SetZoom sets the camera zoom factor
func (cam *Camera) SetZoom() int {
	return cam.zoom
}

// FOV returns the current camera field of view
func (cam *Camera) FOV() float32 {
	return cam.fov
}

// SetType sets the type of cmera
func (cam *Camera) SetType(ct CameraType) {
	cam.id = ct
}

// Type returns the type of camera
func (cam *Camera) Type() CameraType{
	return cam.id
}

// New creates a new perspective camera
func New(fov, aspect, near, far float32) *Camera {
	return &Camera{
		id: Perspective, 
		fov: fov, 
		aspect: aspect, 
		near: near, 
		far: far,
	}
}
