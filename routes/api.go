package routes

import (
	"macaron-simple/controllers/api"

	"gopkg.in/macaron.v1"
)

func InitApiRoutes(m *macaron.Macaron) {
	m.Group("/api", func() {
		m.Get("/", api.DefaultController{}.Hello)
	})

}
