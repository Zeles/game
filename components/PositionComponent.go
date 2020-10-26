package components

import (
	"github.com/faiface/pixel"
	"reflect"
)

type PositionComponent struct {
	Position pixel.Vec
	isActive bool
}

func (Position *PositionComponent)Init() {
	Position.isActive = false
}

func (Position *PositionComponent)GetName() string {
	return reflect.TypeOf(Position).String()
}

func (Position *PositionComponent)GetActive() bool {
	return Position.isActive
}

func (Position *PositionComponent)SetActive(active bool) {
	active = true
	Position.isActive = active
}
