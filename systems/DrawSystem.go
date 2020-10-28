package systems

import (
	"game/components"
	"game/components/window"
	"game/ecs"
	"github.com/faiface/pixel"
	"reflect"
)

type DrawSystem struct {
	windowManager *ecs.Entity
	core *ecs.SystemCore
}

func (Draw *DrawSystem)Init(core *ecs.ECSManager) {
	Draw.core = new(ecs.SystemCore)
	Draw.core.IsActive = true
	Draw.core.Core = core
	Draw.core.AddComponentToFilter(&components.DrawComponent{})
}

func (Draw *DrawSystem)Update() {
	Draw.FindWindowManager()
	for _, v := range Draw.core.Entitys {
		tmp := Draw.windowManager.Children
		if v.GetActive() && tmp != nil {
			drawComp := v.GetComponent(&components.DrawComponent{}).(*components.DrawComponent)
			if drawComp.Load {
				for tmp != nil {
					winCom := tmp.GetComponent(&window.WindowComponent{}).(*window.WindowComponent)
					if winCom.WindowId == drawComp.WindowId {
						drawComp.Sprite.Draw(winCom.Window, pixel.IM.Moved(winCom.Window.Bounds().Center()))
						break
					}
					tmp = tmp.Next
				}
			}
		}
	}
}

func (Draw *DrawSystem)GetActive() bool {
	return Draw.core.IsActive
}

func (Draw *DrawSystem)SetActive(active bool) {
	Draw.core.IsActive = active
}

func (Draw *DrawSystem)GetCore() *ecs.SystemCore {
	return Draw.core
}

func (Draw *DrawSystem)GetName() string {
	return reflect.TypeOf(Draw).String()
}

func (Draw *DrawSystem)FindWindowManager() {
	if Draw.windowManager == nil {
		tmp := Draw.core.Core.Entitys
		for tmp != nil && Draw.windowManager == nil {
			if tmp != nil && tmp.HasComponent(&window.WindowManagerComponent{}) {
				Draw.windowManager = tmp
				tmp = nil
				break
			}
			tmp = tmp.Next
		}
	}
}
