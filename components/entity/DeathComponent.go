package entity

import "reflect"

type DeathComponent struct {
	isActive bool
}

func (Death *DeathComponent)Init() {
	Death.isActive = true
}

func (Death *DeathComponent)GetName() string {
	return reflect.TypeOf(Death).String()
}

func (Death *DeathComponent)GetActive() bool {
	return Death.isActive
}

func (Death *DeathComponent)SetActive(active bool) {
	Death.isActive = active
}
