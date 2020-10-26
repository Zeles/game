package scene

import "reflect"

type SceneManagerComponent struct {
	isActive bool
}

func (SceneManager *SceneManagerComponent)Init() {
	SceneManager.isActive = false
}

func (SceneManager *SceneManagerComponent)GetName() string {
	return reflect.TypeOf(SceneManager).String()
}

func (SceneManager *SceneManagerComponent)GetActive() bool {
	return SceneManager.isActive
}

func (SceneManager *SceneManagerComponent)SetActive(active bool) {
	active = true
	SceneManager.isActive = active
}
