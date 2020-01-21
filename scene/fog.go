package scene

import (
	"github.com/spatocode/karid"
)

type Fog struct {
	color karid.Color
}

func (fg *Fog) Color() karid.Color {
	return fg.color
}

func (fg *Fog) SetColor(color karid.Color) {
	fg.color = color
}

func CreateFog(color karid.Color) *Fog {
	return &Fog{}
}