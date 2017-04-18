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

type HelloResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Hello  `json:"data"`
}

func (d DefaultController) Hello(ctx *macaron.Context) {
	r := HelloResult{200, "", Hello{"say hi"}}

	ctx.JSON(200, &r)
}
