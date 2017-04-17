package routes

import (
	"macaron-simple/controllers"

	"gopkg.in/macaron.v1"
)

func InitWebRoutes(m *macaron.Macaron) {

	m.Get("/", controllers.DefaultController{}.Hello)

}
