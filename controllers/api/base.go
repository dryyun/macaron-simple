package api

import (
	"macaron-simple/controllers"

	"gopkg.in/macaron.v1"
)

type BaseApiController struct {
	controllers.BaseController
}

type Result struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NilArray struct{}

func (b BaseApiController) Success(data interface{}) *Result {
	return &Result{Status: 200, Message: "", Data: data}
}

func (b BaseApiController) Failed(status int, message string) *Result {
	return &Result{status, message, NilArray{}}
}

func (b BaseApiController) Hello(ctx *macaron.Context) {
	r := b.Success(NilArray{})
	ctx.JSON(200, &r)
}
