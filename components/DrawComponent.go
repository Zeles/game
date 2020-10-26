package components

import (
	"github.com/faiface/pixel"
	"reflect"
)

type DrawComponent struct {
	Name string
	Load bool
	Sprite *pixel.Sprite
	isActive bool
}

func (Draw *DrawComponent)Init() {
	Draw.isActive = false
}

func (Draw *DrawComponent)GetName() string {
	return reflect.TypeOf(Draw).String()
}

func (Draw *DrawComponent)GetActive() bool {
	return Draw.isActive
}

func (Draw *DrawComponent)SetActive(active bool) {
	Draw.isActive = active
}
