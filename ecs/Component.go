package ecs

type Component interface {
	Init()
	GetName() string
	GetActive() bool
	SetActive(bool)
}