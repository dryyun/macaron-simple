package routes

import "gopkg.in/macaron.v1"

func InitRoutes(m *macaron.Macaron) {
	InitWebRoutes(m)
	InitApiRoutes(m)
}
