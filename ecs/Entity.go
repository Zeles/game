package ecs

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"io/ioutil"
)

type Entity struct {
	core *ECSManager
	isActive bool
	components map[string]Component
	Parent *Entity
	Prev *Entity
	Next *Entity
	Children *Entity
}

func createEntity(core *ECSManager) *Entity {
	entity := new(Entity)
	entity.components = make(map[string]Component)
	entity.core = core
	entity.Init()
	return entity
}

func (root *Entity)NewEntity() *Entity {
	entity := createEntity(root.core)
	if root == nil {
		root = entity
	}
	if root.Next == nil {
		root.Next = entity
		entity.Prev = root
	} else {
		tmp := root
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = entity
		entity.Prev = tmp.Next
	}
	return entity
}

func (root *Entity)NewChildrenEntity() *Entity {
	var entity *Entity
	if root.Children == nil {
		entity = createEntity(root.core)
		entity.Parent = root
		if root.Children == nil {
			root.Children = entity
		} else {
			tmp := root.Children
			for tmp.Next != nil {
				tmp = tmp.Next
			}
			entity.Prev = tmp.Next
			tmp.Next = entity
		}
	} else {
		entity = root.Children.NewEntity()
	}
	return entity
}

func (entity *Entity)Init() {
	entity.components = make(map[string]Component)
}

func (entity *Entity)AddComponent(component Component) {
	entity.components[component.GetName()] = component
	entity.components[component.GetName()].Init()
	for _, system := range entity.core.systems {
		filtered := true
		core := system.GetCore()
		for _, fc := range core.Components {
			if !entity.HasComponent(fc) {
				filtered = false
				break
			}
		}
		if filtered {
			core.addEntity(entity)
		}
	}
}

func (entity *Entity)RemoveComponent(component Component) {
	delete(entity.components, component.GetName())
	for _, system := range entity.core.systems {
		filtered := true
		core := system.GetCore()
		for _, fc := range core.Components {
			if !entity.HasComponent(fc) {
				filtered = false
				break
			}
		}
		if !filtered {
			core.removeEntity(entity)
		}
	}
}

func (entity *Entity)GetComponent(component Component) Component {
	return entity.components[component.GetName()]
}

func (entity *Entity)HasComponent(component Component) bool {
	if entity.components[component.GetName()] != nil {
		return true
	}
	return false
}

func (entity *Entity)GetActive() bool {
	return entity.isActive
}

func (entity *Entity)SetActive(active bool) {
	entity.isActive = active
}

func (entity *Entity)GetRecursiveUpActive() bool {
	if entity.isActive {
		if entity.Parent == nil {
			return true
		}
		return entity.Parent.GetRecursiveUpActive()
	}
	return false
}

func (entity *Entity)GetRecursiveDownActive() bool {
	if entity.isActive {
		if entity.Children == nil {
			return true
		}
		return entity.Children.GetRecursiveDownActive()
	}
	return false
}

func (entity *Entity)SavePrefub(filepath string) error {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(entity)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, []byte(base64.StdEncoding.EncodeToString(b.Bytes())), 0644)
}
