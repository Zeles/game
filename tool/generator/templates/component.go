package templates

var ComponentTemplate = `package {{.Package}}

import "reflect"

type {{.Name}}Component struct {
	isActive bool
}

func ({{.Name}} *{{.Name}}Component)Init() {
	{{.Name}}.isActive = false
}

func ({{.Name}} *{{.Name}}Component)GetName() string {
	return reflect.TypeOf({{.Name}}).String()
}

func ({{.Name}} *{{.Name}}Component)GetActive() bool {
	return {{.Name}}.isActive
}

func ({{.Name}} *{{.Name}}Component)SetActive(active bool) {
	{{.Name}}.isActive = active
}
`
