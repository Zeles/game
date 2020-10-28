package player

import (
	"github.com/faiface/pixel"
	"reflect"
)

type PlayerMovingComponent struct {
	MovingPosition pixel.Vec
	isActive bool
}

func (PlayerMoving *PlayerMovingComponent)Init() {
	PlayerMoving.isActive = true
}

func (PlayerMoving *PlayerMovingComponent)GetName() string {
	return reflect.TypeOf(PlayerMoving).String()
}

func (PlayerMoving *PlayerMovingComponent)GetActive() bool {
	return PlayerMoving.isActive
}

func (PlayerMoving *PlayerMovingComponent)SetActive(active bool) {
	PlayerMoving.isActive = active
}
