package main

import (
	"game/ecs"
	"game/systems"
	"game/systems/asset"
	"game/systems/scene"
	"game/systems/window"
)

func LoadSystem(manager *ecs.ECSManager) {
	manager.NewSystem(&window.InputSystem{})

	manager.NewSystem(&asset.LoadImageSystem{})
	manager.NewSystem(&asset.LoadSoundSystem{})
	manager.NewSystem(&asset.LoadSpriteSystem{})

	manager.NewSystem(&scene.ChangeSceneSystem{})

	manager.NewSystem(&systems.DrawSystem{})
}
