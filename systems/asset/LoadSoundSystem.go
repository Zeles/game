package asset

import (
	"game/components/asset"
	"game/ecs"
	"game/utils"
	"log"
	"reflect"
)

type LoadSoundSystem struct {
	core *ecs.SystemCore
}

func (LoadSound *LoadSoundSystem)Init(core *ecs.ECSManager) {
	LoadSound.core = new(ecs.SystemCore)
	LoadSound.core.Core = core
	LoadSound.core.IsActive = true
	LoadSound.core.AddComponentToFilter(&asset.SoundComponent{})
}

func (LoadSound *LoadSoundSystem)Update() {
	var err error
	for _, v := range LoadSound.core.Entitys {
		snd := v.GetComponent(&asset.SoundComponent{}).(*asset.SoundComponent)
		if !snd.Load {
			snd.Stream, snd.Format, err = utils.LoadSound(snd.Path, snd.Mode)
			if err != nil {
				log.Panic(err)
			}
			snd.Load = false
		}
	}
	LoadSound.SetActive(false)
}

func (LoadSound *LoadSoundSystem)GetActive() bool {
	return LoadSound.core.IsActive
}

func (LoadSound *LoadSoundSystem)SetActive(active bool) {
	LoadSound.core.IsActive = active
}

func (LoadSound *LoadSoundSystem)GetCore() *ecs.SystemCore {
	return LoadSound.core
}

func (LoadSound *LoadSoundSystem)GetName() string {
	return reflect.TypeOf(LoadSound).String()
}
