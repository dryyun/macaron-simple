package controllers

type Result struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NilArray struct {
}

type BaseController struct {
}

func (b BaseController) Success(data interface{}) *Result {
	return &Result{Status: 200, Message: "", Data: data}
}

func (b BaseController) Failed(status int, message string) *Result {
	return &Result{status, message, NilArray{}}
}
