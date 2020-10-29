package scenes

import (
	sceneComp "game/components/scene"
	"game/ecs"
)

func MainMenu(root *ecs.Entity) *ecs.Entity {
	mainMenuScene := root.NewChildrenEntity()
	mainMenuComp := sceneComp.SceneComponent{Name: "MainMenu"}
	mainMenuScene.SetActive(true)
	mainMenuScene.AddComponent(&mainMenuComp)
	return mainMenuScene
}
