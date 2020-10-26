package scene

import "reflect"

type SceneComponent struct {
	Name string
	isActive bool
}

func (scene *SceneComponent)Init() {
	scene.isActive = false
}

func (scene *SceneComponent)GetName() string {
	return reflect.TypeOf(scene).String()
}

func (scene *SceneComponent)GetActive() bool {
	return scene.isActive
}

func (scene *SceneComponent)SetActive(active bool) {
	scene.isActive = active
}
