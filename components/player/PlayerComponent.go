package player

import "reflect"

type PlayerComponent struct {
	isActive bool
}

func (Player *PlayerComponent)Init() {
	Player.isActive = true
}

func (Player *PlayerComponent)GetName() string {
	return reflect.TypeOf(Player).String()
}

func (Player *PlayerComponent)GetActive() bool {
	return Player.isActive
}

func (Player *PlayerComponent)SetActive(active bool) {
	Player.isActive = active
}
