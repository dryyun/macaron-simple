package api

import (
	"macaron-simple/controllers"

	macaron "gopkg.in/macaron.v1"
)

type DefaultController struct {
	controllers.BaseController
}

type Hello struct {
	Action string `json:"action"`
}

func (d DefaultController) Hello(ctx *macaron.Context) {

	r := d.Success(Hello{"say hi"})

	ctx.JSON(200, &r)
}
