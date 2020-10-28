package window

import "reflect"

type WindowManagerComponent struct {
	isActive bool
}

func (WindowManager *WindowManagerComponent)Init() {
	WindowManager.isActive = true
}

func (WindowManager *WindowManagerComponent)GetName() string {
	return reflect.TypeOf(WindowManager).String()
}

func (WindowManager *WindowManagerComponent)GetActive() bool {
	return WindowManager.isActive
}

func (WindowManager *WindowManagerComponent)SetActive(active bool) {
	WindowManager.isActive = active
}
