package main

import (
	"game/components/asset"
	"game/ecs"
)

func LoadAssets(root *ecs.Entity) *ecs.Entity {
	assetManager := root.NewEntity()
	assetManager.AddComponent(&asset.AssetManagerComponent{})

	img := root.NewChildrenEntity()
	img.AddComponent(&asset.ImageComponent{Name: "ascii", Path: "./resorces/font/ascii.png"})
	return img
}
