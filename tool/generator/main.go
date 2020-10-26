package main

import (
	"flag"
	"game/tool/generator/templates"
	"html/template"
	"log"
	"os"
	"strings"
)

type Mode int32

const (
	component Mode = 0
	system Mode = 1
)

func GetMode(mode string) Mode {
	var r Mode
	mode = strings.ToLower(mode)
	switch mode {
	case "c", "component":
		r = component
	case "s", "system":
		r = system
	}
	return r
}

type Data struct {
	Mode Mode
	Name string
	Package string
}

func GetPackageName(path string) string {
	count := len(path) - 1
	for count > 0 {
		if path[count - 1] == '/' {
			return path[count:]
		}
		count--
	}
	if path == "." {
		return "main"
	}
	return path
}

func main() {
	var d Data
	var path, mode, newfilepath string
	var t *template.Template
	flag.StringVar(&mode, "mode", "component", "The mode generation [component, system, etd]")
	flag.StringVar(&d.Name,"name", "test", "The name")
	flag.StringVar(&path, "path", ".", "The create path")
	flag.Parse()
	d.Mode = GetMode(mode)
	d.Package = GetPackageName(path)

	switch d.Mode {
	case component:
		t = template.Must(template.New("component").Parse(templates.ComponentTemplate))
		newfilepath = path + "/" + d.Name + "Component.go"
	case system:
		t = template.Must(template.New("system").Parse(templates.SystemTemplate))
		newfilepath = path + "/" + d.Name + "System.go"
	}

	file, err := os.Create(newfilepath)
	if err != nil {
		log.Panic(err)
	}

	t.Execute(file, d)
}
