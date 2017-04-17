package controllers

import macaron "gopkg.in/macaron.v1"

type DefaultController struct {
}

func (d DefaultController) Hello(ctx *macaron.Context) {
	ctx.Data["Name"] = "hello"
	ctx.HTML(200, "default/hello") // 200 为响应码
}
