package main

import (
	"game/components"
	sceneComp "game/components/scene"
	"game/components/window"
	"game/ecs"
	"game/scenes"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

func run() {
	manager := ecs.NewECSManager()
	LoadSystem(manager)

	windowManager := manager.Entitys.NewEntity()
	windowManager.AddComponent(&window.WindowManagerComponent{})
	mainWindow := windowManager.NewChildrenEntity()
	mainWindow.AddComponent(&window.WindowComponent{WindowId: "main", Config: pixelgl.WindowConfig{
		Title: "main",
		Bounds: pixel.Rect{pixel.Vec{0, 0}, pixel.Vec{800, 600}}}})
	mainWindow.AddComponent(&window.InputManagerComponent{})

	sceneManager := manager.Entitys.NewEntity()
	sceneManager.AddComponent(&sceneComp.SceneManagerComponent{})

	mainMenuScene := scenes.MainMenu(sceneManager)
	gameScene := scenes.Game(sceneManager)

	gameScene.GetActive()

	LoadAssets(manager.Entitys)

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
