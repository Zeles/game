package systems

import (
	"game/components"
	"game/ecs"
	"github.com/faiface/pixel"
	"reflect"
)

type DrawSystem struct {
	core *ecs.SystemCore
}

func (Draw *DrawSystem)Init(core *ecs.ECSManager) {
	Draw.core = new(ecs.SystemCore)
	Draw.core.IsActive = true
	Draw.core.Core = core
	Draw.core.AddComponentToFilter(&components.DrawComponent{})
}

func (Draw *DrawSystem)Update() {
	for _, v := range Draw.core.Entitys {
		drawComp := v.GetComponent(&components.DrawComponent{}).(*components.DrawComponent)
		if drawComp.Load {
			drawComp.Sprite.Draw(Draw.core.Core.Win, pixel.IM.Moved(Draw.core.Core.Win.Bounds().Center()))
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
