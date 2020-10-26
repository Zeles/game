package ecs

type SystemCore struct {
	Core *ECSManager
	IsActive bool
	Components []Component
	Entitys []*Entity
}

func (systemCore *SystemCore)AddComponentToFilter(components ...Component) {
	for _, v := range components {
		systemCore.Components = append(systemCore.Components, v)
	}
}

func (systemCore *SystemCore)addEntity(entity *Entity) {
	if len(systemCore.Components) > 0 {
		for _, v := range systemCore.Entitys {
			if entity == v {
				return
			}
		}
		systemCore.Entitys = append(systemCore.Entitys, entity)
	}
}

func (systemCore *SystemCore)removeEntity(entity *Entity) {
	for k, v := range systemCore.Entitys {
		if v == entity {
			systemCore.Entitys[k] = systemCore.Entitys[len(systemCore.Entitys) - 1]
			systemCore.Entitys[len(systemCore.Entitys) - 1] = nil
			systemCore.Entitys = systemCore.Entitys[:len(systemCore.Entitys) - 1]
			break
		}
	}
}

type System interface {
	Init(*ECSManager)
	Update()
	GetActive() bool
	SetActive(bool)
	GetCore() *SystemCore
	GetName() string
}
