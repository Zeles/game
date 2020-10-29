package window

import (
	"game/components/window"
	"game/ecs"
	"reflect"
)

type InputSystem struct {
	core *ecs.SystemCore
}

func (Input *InputSystem)Init(core *ecs.ECSManager) {
	Input.core = new(ecs.SystemCore)
	Input.core.Core = core
	Input.core.IsActive = true
	Input.core.AddComponentToFilter(&window.WindowComponent{}, &window.InputManagerComponent{})
}

func (Input *InputSystem)Update() {
	for _, v := range Input.core.Entitys {
		if v.GetActive() {
			input := v.GetComponent(&window.InputManagerComponent{}).(*window.InputManagerComponent)
			window := v.GetComponent(&window.WindowComponent{}).(*window.WindowComponent)
			input.MousePos = window.Window.MousePosition()
			for key, button := range input.ButtonIds {
				for _, b := range button {
					if !input.Buttons[key] {
						if input.Buttons[key] = window.Window.Pressed(b); input.Buttons[key] {
							break
						}
					}
				}
				input.Buttons[key] = false
			}
		}
	}
}

func (Input *InputSystem)GetActive() bool {
	return Input.core.IsActive
}

func (Input *InputSystem)SetActive(active bool) {
	Input.core.IsActive = active
}

func (Input *InputSystem)GetCore() *ecs.SystemCore {
	return Input.core
}

func (Input *InputSystem)GetName() string {
	return reflect.TypeOf(Input).String()
}
