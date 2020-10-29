package scenes

import (
	sceneComp "game/components/scene"
	"game/ecs"
)

func Game(root *ecs.Entity) *ecs.Entity {
	gameScene := root.NewChildrenEntity()
	gameComp := sceneComp.SceneComponent{Name: "Game"}
	gameComp.SetActive(false)
	gameScene.AddComponent(&gameComp)
	gameScene.AddComponent(&sceneComp.ChangeSceneComponent{})
	return gameScene
}
