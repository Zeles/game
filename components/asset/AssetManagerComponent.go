package asset

import "reflect"

type AssetManagerComponent struct {
	isActive bool
}

func (AssetManager *AssetManagerComponent)Init() {
	AssetManager.isActive = false
}

func (AssetManager *AssetManagerComponent)GetName() string {
	return reflect.TypeOf(AssetManager).String()
}

func (AssetManager *AssetManagerComponent)GetActive() bool {
	return AssetManager.isActive
}

func (AssetManager *AssetManagerComponent)SetActive(active bool) {
	active = true
	AssetManager.isActive = active
}
