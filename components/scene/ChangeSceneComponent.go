package scene

import "reflect"

type ChangeSceneComponent struct {
	isActive bool
}

func (ChangeScene *ChangeSceneComponent)Init() {
	ChangeScene.isActive = false
}

func (ChangeScene *ChangeSceneComponent)GetName() string {
	return reflect.TypeOf(ChangeScene).String()
}

func (ChangeScene *ChangeSceneComponent)GetActive() bool {
	return ChangeScene.isActive
}

func (ChangeScene *ChangeSceneComponent)SetActive(active bool) {
	ChangeScene.isActive = active
}
