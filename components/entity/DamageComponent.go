package entity

import "reflect"

type DamageComponent struct {
	Count int32
	isActive bool
}

func (Damage *DamageComponent)Init() {
	Damage.isActive = true
}

func (Damage *DamageComponent)GetName() string {
	return reflect.TypeOf(Damage).String()
}

func (Damage *DamageComponent)GetActive() bool {
	return Damage.isActive
}

func (Damage *DamageComponent)SetActive(active bool) {
	Damage.isActive = active
}
