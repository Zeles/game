package components

import (
	"github.com/faiface/pixel"
	"reflect"
)

type ScaleComponent struct {
	Scale pixel.Vec
	isActive bool
}

func (Scale *ScaleComponent)Init() {
	Scale.isActive = false
}

func (Scale *ScaleComponent)GetName() string {
	return reflect.TypeOf(Scale).String()
}

func (Scale *ScaleComponent)GetActive() bool {
	return Scale.isActive
}

func (Scale *ScaleComponent)SetActive(active bool) {
	active = true
	Scale.isActive = active
}
