package asset

import (
	"game/components/asset"
	"game/ecs"
	"game/utils"
	"log"
	"reflect"
)

type LoadImageSystem struct {
	core *ecs.SystemCore
}

func (LoadImage *LoadImageSystem)Init(core *ecs.ECSManager) {
	LoadImage.core = new(ecs.SystemCore)
	LoadImage.core.Core = core
	LoadImage.core.IsActive = true
	LoadImage.core.AddComponentToFilter(&asset.ImageComponent{})
}

func (LoadImage *LoadImageSystem)Update() {
	var err error
	for _, v := range LoadImage.core.Entitys {
		img := v.GetComponent(&asset.ImageComponent{}).(*asset.ImageComponent)
		if !img.Load {
			img.Image, err = utils.LoadPicture(img.Path)
			if err != nil {
				log.Panic(err)
			}
			img.Load = true
		}
	}
	LoadImage.SetActive(false)
}

func (LoadImage *LoadImageSystem)GetActive() bool {
	return LoadImage.core.IsActive
}

func (LoadImage *LoadImageSystem)SetActive(active bool) {
	LoadImage.core.IsActive = active
}

func (LoadImage *LoadImageSystem)GetCore() *ecs.SystemCore {
	return LoadImage.core
}

func (LoadImage *LoadImageSystem)GetName() string {
	return reflect.TypeOf(LoadImage).String()
}
