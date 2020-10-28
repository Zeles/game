package entity

import "reflect"

type HealthComponent struct {
	Max int32
	Current int32
	isActive bool
}

func (Health *HealthComponent)Init() {
	Health.isActive = true
	Health.Current = Health.Max
}

func (Health *HealthComponent)GetName() string {
	return reflect.TypeOf(Health).String()
}

func (Health *HealthComponent)GetActive() bool {
	return Health.isActive
}

func (Health *HealthComponent)SetActive(active bool) {
	Health.isActive = active
}
