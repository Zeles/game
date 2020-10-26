package main

import (
	"game/components"
	"game/components/asset"
	sceneComp "game/components/scene"
	"game/ecs"
	"game/systems"
	asset2 "game/systems/asset"
	sceneSys "game/systems/scene"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"log"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:                  "Game",
		Bounds:                 pixel.Rect{pixel.Vec{X: 0, Y: 0}, pixel.Vec{X: 800, Y: 600}},
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Panic(err)
	}
	manager := ecs.NewECSManager()
	manager.Win = win
	manager.NewSystem(&asset2.LoadImageSystem{})
	manager.NewSystem(&asset2.LoadSoundSystem{})
	manager.NewSystem(&asset2.LoadSpriteSystem{})
	manager.NewSystem(&sceneSys.ChangeSceneSystem{})
	manager.NewSystem(&systems.DrawSystem{})

	sceneManager := manager.Entitys.NewEntity()
	sceneManager.AddComponent(&sceneComp.SceneManagerComponent{})
	mainMenuScene := sceneManager.NewChildrenEntity()
	mainMenuComp := sceneComp.SceneComponent{Name: "MainMenu"}
	mainMenuScene.SetActive(true)
	mainMenuScene.AddComponent(&mainMenuComp)
	gameScene := sceneManager.NewChildrenEntity()
	gameComp := sceneComp.SceneComponent{Name: "Game"}
	gameComp.SetActive(false)
	gameScene.AddComponent(&gameComp)
	gameScene.AddComponent(&sceneComp.ChangeSceneComponent{})

	assetManager := manager.Entitys.NewEntity()
	assetManager.AddComponent(&asset.AssetManagerComponent{})

	img := assetManager.NewChildrenEntity()
	img.AddComponent(&asset.ImageComponent{Name: "ascii", Path: "./resorces/font/ascii.png"})

	draw := mainMenuScene.NewEntity()
	draw.AddComponent(&components.DrawComponent{Name: "ascii"})

	for !win.Closed() {
		win.Clear(color.Black)
		manager.UpdateSystem()
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
