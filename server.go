package main

import (
	"html/template"
	"macaron-simple/routes"

	"github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Use(gzip.Gziper())
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  "views",
		Extensions: []string{".html"},
		Funcs: []template.FuncMap{map[string]interface{}{
			"AppName": func() string {
				return "Macaron"
			},
		}},
	}))

	routes.InitRoutes(m)

	m.Run()
}
