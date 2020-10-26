package templates

var SystemTemplate = `package {{.Package}}

import (
	"game/ecs"
	"reflect"
)

type {{.Name}}System struct {
	core *ecs.SystemCore
}

func ({{.Name}} *{{.Name}}System)Init(core *ecs.ECSManager) {
	{{.Name}}.core = new(ecs.SystemCore)
	{{.Name}}.core.Core = core
	{{.Name}}.core.IsActive = true
}

func ({{.Name}} *{{.Name}}System)Update() {
	for _, v := range {{.Name}}.core.Entitys {

	}
}

func ({{.Name}} *{{.Name}}System)GetActive() bool {
	return {{.Name}}.core.IsActive
}

func ({{.Name}} *{{.Name}}System)SetActive(active bool) {
	{{.Name}}.core.IsActive = active
}

func ({{.Name}} *{{.Name}}System)GetCore() *ecs.SystemCore {
	return {{.Name}}.core
}

func ({{.Name}} *{{.Name}}System)GetName() string {
	return reflect.TypeOf({{.Name}}).String()
}
`
