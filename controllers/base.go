package controllers

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type BaseController struct {
}

func (b BaseController) Success() *Result {
	return &Result{200, ""}
}

func (b BaseController) Failed(status int, message string) *Result {
	return &Result{status, message}
}
