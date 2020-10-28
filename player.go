package main

import (
	"game/components"
	"game/components/entity"
	"game/components/player"
	"game/ecs"
)

func CreatePlayer(root *ecs.Entity) *ecs.Entity {
	ePlayer := root.NewEntity()
	ePlayer.AddComponent(&player.PlayerComponent{})
	ePlayer.AddComponent(&player.PlayerMovingComponent{})
	ePlayer.AddComponent(&entity.HealthComponent{Max: 100})
	ePlayer.AddComponent(&components.DrawComponent{WindowId:  "main", Name: ""})
	return ePlayer
}
