package utils

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	"os"
	_ "github.com/faiface/beep/mp3"
	_ "github.com/faiface/beep/wav"
	_ "github.com/faiface/beep/flac"
)

type SoundFormat int32

const (
	fsMp3 SoundFormat = 0
	fsWav SoundFormat = 1
	fsFlac SoundFormat = 2
)

func LoadSound(path string, format SoundFormat) (beep.StreamSeekCloser, beep.Format, error) {
	var stream beep.StreamSeekCloser
	var form beep.Format

	file, err := os.Open(path)
	if err != nil {
		return nil, form, err
	}
	switch format {
	case fsMp3:
		stream, form, err = mp3.Decode(file)
	case fsWav:
		stream, form, err = wav.Decode(file)
	case fsFlac:
		stream, form, err = flac.Decode(file)
	}
	return stream, form, nil
}
