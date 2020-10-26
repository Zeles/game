package asset

import (
	"game/utils"
	"github.com/faiface/beep"
	"log"
	"reflect"
)

type SoundComponent struct {
	Path string
	Load bool
	Mode utils.SoundFormat
	Stream beep.StreamSeekCloser
	Format beep.Format
	isActive bool
}

func (Sound *SoundComponent)Init() {
	var err error
	Sound.isActive = false
	Sound.Stream, Sound.Format, err = utils.LoadSound(Sound.Path, Sound.Mode)
	if err != nil {
		log.Panic(err)
	}
}

func (Sound *SoundComponent)GetName() string {
	return reflect.TypeOf(Sound).String()
}

func (Sound *SoundComponent)GetActive() bool {
	return Sound.isActive
}

func (Sound *SoundComponent)SetActive(active bool) {
	Sound.isActive = active
}
