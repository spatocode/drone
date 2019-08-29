package drivers

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/spatocode/drone"
)

func keyboardTranslate(key glfw.Key) drone.KeyboardKey {
	switch key {
	case glfw.Key0:
		return drone.Key0
	case glfw.Key1:
		return drone.Key1
	case glfw.Key2:
		return drone.Key2
	case glfw.Key3:
		return drone.Key3
	case glfw.Key4:
		return drone.Key4
	case glfw.Key5:
		return drone.Key5
	case glfw.Key6:
		return drone.Key6
	case glfw.Key7:
		return drone.Key7
	case glfw.Key8:
		return drone.Key8
	case glfw.Key9:
		return drone.Key9
	default:
		return ""
	}
}
