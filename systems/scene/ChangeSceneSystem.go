package scene

import (
	"game/components/scene"
	"game/ecs"
	"reflect"
)

type ChangeSceneSystem struct {
	core *ecs.SystemCore
}

func (ChangeScene *ChangeSceneSystem)Init(core *ecs.ECSManager) {
	ChangeScene.core = new(ecs.SystemCore)
	ChangeScene.core.Core = core
	ChangeScene.core.IsActive = true
	ChangeScene.core.AddComponentToFilter(&scene.SceneComponent{})
}

func (ChangeScene *ChangeSceneSystem)Update() {
	for _, v := range ChangeScene.core.Entitys {
		if v.HasComponent(&scene.ChangeSceneComponent{}) {
			v.SetActive(true)
			v.RemoveComponent(&scene.ChangeSceneComponent{})
		} else {
			v.SetActive(false)
		}
	}
}

func (ChangeScene *ChangeSceneSystem)GetActive() bool {
	return ChangeScene.core.IsActive
}

func (ChangeScene *ChangeSceneSystem)SetActive(active bool) {
	ChangeScene.core.IsActive = active
}

func (ChangeScene *ChangeSceneSystem)GetCore() *ecs.SystemCore {
	return ChangeScene.core
}

func (ChangeScene *ChangeSceneSystem)GetName() string {
	return reflect.TypeOf(ChangeScene).String()
}
