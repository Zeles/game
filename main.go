package main

import (
	"game/components"
	"game/components/asset"
	sceneComp "game/components/scene"
	"game/components/window"
	"game/ecs"
	"game/systems"
	asset2 "game/systems/asset"
	sceneSys "game/systems/scene"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

func run() {
	manager := ecs.NewECSManager()
	manager.NewSystem(&asset2.LoadImageSystem{})
	manager.NewSystem(&asset2.LoadSoundSystem{})
	manager.NewSystem(&asset2.LoadSpriteSystem{})
	manager.NewSystem(&sceneSys.ChangeSceneSystem{})
	manager.NewSystem(&systems.DrawSystem{})

	windowManager := manager.Entitys.NewEntity()
	windowManager.AddComponent(&window.WindowManagerComponent{})
	mainWindow := windowManager.NewChildrenEntity()
	mainWindow.AddComponent(&window.WindowComponent{WindowId: "main", Config: pixelgl.WindowConfig{
		Title: "main",
		Bounds: pixel.Rect{pixel.Vec{0, 0}, pixel.Vec{800, 600}}}})

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
	draw.AddComponent(&components.DrawComponent{Name: "ascii", WindowId: "main"})

	mainWin := mainWindow.GetComponent(&window.WindowComponent{}).(*window.WindowComponent)

	for !mainWin.Window.Closed() {
		mainWin.Window.Clear(color.Black)
		manager.UpdateSystem()
		mainWin.Window.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
