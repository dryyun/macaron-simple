package main

import (
	"html/template"
	"macaron-simple/routes"

	"github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(gzip.Gziper())
	m.Use(macaron.Static("public"))

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
