package gui

import (
	"game/ecs"
	"reflect"
)

type ClickSystem struct {
	core *ecs.SystemCore
}

func (Click *ClickSystem)Init(core *ecs.ECSManager) {
	Click.core = new(ecs.SystemCore)
	Click.core.Core = core
	Click.core.IsActive = true
}

func (Click *ClickSystem)Update() {
	for _, v := range Click.core.Entitys {
		if v.GetActive() {

		}
	}
}

func (Click *ClickSystem)GetActive() bool {
	return Click.core.IsActive
}

func (Click *ClickSystem)SetActive(active bool) {
	Click.core.IsActive = active
}

func (Click *ClickSystem)GetCore() *ecs.SystemCore {
	return Click.core
}

func (Click *ClickSystem)GetName() string {
	return reflect.TypeOf(Click).String()
}
