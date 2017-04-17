package api

import macaron "gopkg.in/macaron.v1"

type DefaultController struct {
}

func (d DefaultController) Hello(ctx *macaron.Context) string {
	return "default api hello"
}
