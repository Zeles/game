package ecs

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"io/ioutil"
	"os"
)

type ECSManager struct {
	Entitys *Entity
	systems map[string]System
	currentEntityId int64
}

func NewECSManager() *ECSManager {
	ecsManager := new(ECSManager)
	ecsManager.systems = make(map[string]System)

	entity := new(Entity)
	entity.components = make(map[string]Component)
	entity.core = ecsManager
	entity.Init()
	ecsManager.Entitys = entity

	return ecsManager
}

func (ecsManager *ECSManager)ComponentRegister(components ...Component) {
	for _, v := range components {
		gob.Register(v)
	}
}

func (ecsManager *ECSManager)NewSystem(system System) {
	ecsManager.systems[system.GetName()] = system
	ecsManager.systems[system.GetName()].Init(ecsManager)
}

func (ecsManager *ECSManager)DeleteSystem(system System) {
	delete(ecsManager.systems, system.GetName())
}

func (ecsManager *ECSManager)UpdateSystem() {
	for _, v := range ecsManager.systems {
		if v.GetActive() {
			v.Update()
		}
	}
}

func (ecsManager *ECSManager)LoadPrefub(filepath string) (*Entity, error) {
	readEn := Entity{}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	str, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	b := bytes.Buffer{}
	b.Write([]byte(str))
	d := gob.NewDecoder(&b)
	err = d.Decode(&readEn)
	if err != nil {
		return nil, err
	}
	entity := ecsManager.Entitys.NewEntity()
	for _, v := range readEn.components {
		entity.AddComponent(v)
	}
	return entity, nil
}
