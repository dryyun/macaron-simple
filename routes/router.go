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
		m.Get("/", api.DefaultController{}.Hello)
	})
}

func InitWebRoutes(m *macaron.Macaron) {
	m.Get("/", controllers.DefaultController{}.Hello)
}
