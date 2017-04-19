package routes

import (
	"macaron-simple/controllers"
	"macaron-simple/controllers/api"

	"gopkg.in/macaron.v1"
)

func InitRoutes(m *macaron.Macaron) {
	InitWebRoutes(m)
	InitApiRoutes(m)
}

func InitApiRoutes(m *macaron.Macaron) {
	m.Group("/api", func() {
		m.Get("/hello/:name", func(ctx *macaron.Context) string {
			return "hello " + ctx.Params("name")
		})

		m.Get("/", api.DefaultController{}.Hello)
	})
}

func InitWebRoutes(m *macaron.Macaron) {

	m.Get("/hello/:name", func(ctx *macaron.Context) string {
		return "hello " + ctx.Params("name")
	})

	m.Get("/", controllers.DefaultController{}.Hello)
}
