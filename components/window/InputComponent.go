package window

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"reflect"
)

type InputManagerComponent struct {
	ButtonIds map[string][]pixelgl.Button
	Buttons map[string]bool
	MousePos pixel.Vec
	isActive bool
}

func (Input *InputManagerComponent)Init() {
	Input.isActive = true
	Input.ButtonIds = make(map[string][]pixelgl.Button)
	Input.Buttons = make(map[string]bool)

	Input.addButton("Exit", []pixelgl.Button{pixelgl.KeyEscape})
	Input.addButton("Left Move", []pixelgl.Button{pixelgl.KeyLeft, pixelgl.KeyA})
	Input.addButton("Right Move", []pixelgl.Button{pixelgl.KeyRight, pixelgl.KeyD})
	Input.addButton("Up Move", []pixelgl.Button{pixelgl.KeyUp, pixelgl.KeyW})
	Input.addButton("Down Move", []pixelgl.Button{pixelgl.KeyDown, pixelgl.KeyS})
	Input.addButton("Inventory", []pixelgl.Button{pixelgl.KeyI})
	Input.addButton("Character", []pixelgl.Button{pixelgl.KeyC})
}

func (Input *InputManagerComponent)GetName() string {
	return reflect.TypeOf(Input).String()
}

func (Input *InputManagerComponent)GetActive() bool {
	return Input.isActive
}

func (Input *InputManagerComponent)SetActive(active bool) {
	Input.isActive = active
}

func (Input *InputManagerComponent)addButton(name string, button []pixelgl.Button) {
	Input.ButtonIds[name] = append(Input.ButtonIds[name], button...)
	Input.Buttons[name] = false
}
