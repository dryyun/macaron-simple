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
		//For Test
		m.Get("/", api.BaseApiController{}.Hello)

		//For Discussion
		m.Post("/user/register", api.UserApiController{}.Register)
	})
}

func InitWebRoutes(m *macaron.Macaron) {

	//For Test
	m.Get("/hello/:name", func(ctx *macaron.Context) string {
		return "hello " + ctx.Params("name")
	})

	m.Get("/", controllers.DefaultController{}.Hello)

	//For Discussion
}
