package asset

import (
	"game/components"
	"game/components/asset"
	"game/ecs"
	"github.com/faiface/pixel"
	"reflect"
)

type LoadSpriteSystem struct {
	assetManager *ecs.Entity
	core *ecs.SystemCore
}

func (LoadSprite *LoadSpriteSystem)Init(core *ecs.ECSManager) {
	LoadSprite.core = new(ecs.SystemCore)
	LoadSprite.core.Core = core
	LoadSprite.core.IsActive = true
	LoadSprite.core.AddComponentToFilter(&components.DrawComponent{})
}

func (LoadSprite *LoadSpriteSystem)Update() {
	LoadSprite.FindAssetManager()
	if LoadSprite.assetManager != nil {
		for _, v := range LoadSprite.core.Entitys {
			sprt := v.GetComponent(&components.DrawComponent{}).(*components.DrawComponent)
			if !sprt.Load {
				if tmp := LoadSprite.assetManager.Children; tmp != nil {
					for tmp != nil {
						if tmp.HasComponent(&asset.ImageComponent{}) {
							img := tmp.GetComponent(&asset.ImageComponent{}).(*asset.ImageComponent)
							if img.Name == sprt.Name {
								sprt.Sprite = pixel.NewSprite(img.Image, img.Image.Bounds())
								sprt.Load = true
								break
							}
							tmp = tmp.Next
						}
					}
				}
			}
		}
	}
}

func (LoadSprite *LoadSpriteSystem)GetActive() bool {
	return LoadSprite.core.IsActive
}

func (LoadSprite *LoadSpriteSystem)SetActive(active bool) {
	LoadSprite.core.IsActive = active
}

func (LoadSprite *LoadSpriteSystem)GetCore() *ecs.SystemCore {
	return LoadSprite.core
}

func (LoadSprite *LoadSpriteSystem)GetName() string {
	return reflect.TypeOf(LoadSprite).String()
}

func (LoadSprite *LoadSpriteSystem)FindAssetManager() {
	if LoadSprite.assetManager == nil {
		tmp := LoadSprite.core.Core.Entitys
		for tmp != nil && LoadSprite.assetManager == nil {
			if tmp != nil && tmp.HasComponent(&asset.AssetManagerComponent{}) {
				LoadSprite.assetManager = tmp
				tmp = nil
				break
			}
			tmp = tmp.Next
		}
	}
}
