package window

import (
	"github.com/faiface/pixel/pixelgl"
	"log"
	"reflect"
)

type WindowComponent struct {
	WindowId string
	Config pixelgl.WindowConfig
	Window *pixelgl.Window
	isActive bool
}

func (Window *WindowComponent)Init() {
	var err error
	Window.isActive = true
	Window.Window, err = pixelgl.NewWindow(Window.Config)
	if err != nil {
		log.Panic(err)
	}
}

func (Window *WindowComponent)GetName() string {
	return reflect.TypeOf(Window).String()
}

func (Window *WindowComponent)GetActive() bool {
	return Window.isActive
}

func (Window *WindowComponent)SetActive(active bool) {
	Window.isActive = active
}
