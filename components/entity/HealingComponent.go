package entity

import "reflect"

type HealingComponent struct {
	Count int32
	isActive bool
}

func (Healing *HealingComponent)Init() {
	Healing.isActive = true
}

func (Healing *HealingComponent)GetName() string {
	return reflect.TypeOf(Healing).String()
}

func (Healing *HealingComponent)GetActive() bool {
	return Healing.isActive
}

func (Healing *HealingComponent)SetActive(active bool) {
	Healing.isActive = active
}
