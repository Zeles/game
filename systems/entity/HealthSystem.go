package entity

import (
	"game/components/entity"
	"game/ecs"
	"reflect"
)

type HealthSystem struct {
	core *ecs.SystemCore
}

func (Health *HealthSystem)Init(core *ecs.ECSManager) {
	Health.core = new(ecs.SystemCore)
	Health.core.Core = core
	Health.core.IsActive = true
	Health.core.AddComponentToFilter(&entity.HealthComponent{})
}

func (Health *HealthSystem)Update() {
	for _, v := range Health.core.Entitys {
		if v.GetActive() {
			healthCom := v.GetComponent(&entity.HealthComponent{}).(*entity.HealthComponent)
			if v.HasComponent(&entity.HealingComponent{}) {
				healingCom := v.GetComponent(&entity.HealingComponent{}).(*entity.HealingComponent)
				healthCom.Current += healingCom.Count
			}
			if v.HasComponent(&entity.DamageComponent{}) {
				damageCom := v.GetComponent(&entity.DamageComponent{}).(*entity.DamageComponent)
				healthCom.Current -= damageCom.Count
			}
			if healthCom.Current <= 0 {
				v.AddComponent(&entity.DeathComponent{})
			}
		}
	}
}

func (Health *HealthSystem)GetActive() bool {
	return Health.core.IsActive
}

func (Health *HealthSystem)SetActive(active bool) {
	Health.core.IsActive = active
}

func (Health *HealthSystem)GetCore() *ecs.SystemCore {
	return Health.core
}

func (Health *HealthSystem)GetName() string {
	return reflect.TypeOf(Health).String()
}
