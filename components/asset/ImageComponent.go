package asset

import (
	"game/utils"
	"github.com/faiface/pixel"
	"log"
	"reflect"
)

type ImageComponent struct {
	Name string
	Path string
	Load bool
	Image pixel.Picture
	isActive bool
}

func (Image *ImageComponent)Init() {
	var err error
	Image.isActive = false
	Image.Image, err = utils.LoadPicture(Image.Path)
	if err != nil {
		log.Panic(err)
	}
}

func (Image *ImageComponent)GetName() string {
	return reflect.TypeOf(Image).String()
}

func (Image *ImageComponent)GetActive() bool {
	return Image.isActive
}

func (Image *ImageComponent)SetActive(active bool) {
	Image.isActive = active
}
