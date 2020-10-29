package gui

import (
	"github.com/faiface/pixel"
	"reflect"
)

type ClickComponent struct {
	Rect pixel.Rect
	IsClick bool
	isActive bool
}

func (Click *ClickComponent)Init() {
	Click.isActive = true
	Click.IsClick = false
}

func (Click *ClickComponent)GetName() string {
	return reflect.TypeOf(Click).String()
}

func (Click *ClickComponent)GetActive() bool {
	return Click.isActive
}

func (Click *ClickComponent)SetActive(active bool) {
	Click.isActive = active
}
