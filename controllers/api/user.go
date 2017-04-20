package api

import (
	"fmt"
	. "macaron-simple/form"

	"gopkg.in/macaron.v1"
)

type UserApiController struct {
	BaseApiController
}

func (u UserApiController) Register(ctx *macaron.Context) string {
	var user UserRegisterForm
	ctx.Req.ParseForm()
	r := u.Validator(&user, ctx.Req.Form, UserRegisterFormErrorMsg)

	if r != nil {
		ctx.JSON(200, &r)
	}

	fmt.Printf("%#v\n", user)
	return "!"
}
