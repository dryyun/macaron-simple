package api

import (
	"fmt"
	"macaron-simple/controllers"
	"net/url"

	"reflect"

	"github.com/go-playground/form"
	"gopkg.in/go-playground/validator.v9"
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

type FormData struct {
	Form interface{} `json:"form"`
}

type NilArray struct{}

func (b BaseApiController) Success(data interface{}) *Result {
	return &Result{Status: 0, Message: "", Data: data}
}

func (b BaseApiController) Failed(status int, message string) *Result {
	return &Result{status, message, NilArray{}}
}

func (b BaseApiController) FormFailed(status int, message string, data interface{}) *Result {
	return &Result{Status: status, Message: message, Data: FormData{Form: data}}
}

func (b BaseApiController) Validator(v interface{}, values url.Values, errMsg map[string]string) interface{} {
	decoder := form.NewDecoder()
	err := decoder.Decode(v, values)
	if err != nil {
		return b.Failed(1, "绑定数据失败")
	}

	validate := validator.New()
	err = validate.Struct(v)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return b.Failed(1, "绑定数据失败")
			//fmt.Println(err)
		}

		var temp map[string]string
		temp = make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {

			key := err.Field() + "." + err.Tag()
			if value, ok := errMsg[key]; ok {
				temp[err.Field()] = value
			}
			fmt.Println(temp)
			rrr := reflect.New(reflect.TypeOf(v))

			rrr.FieldByName(err.Field()).Set(reflect.ValueOf("123"))
			fmt.Println(rrr)

			//fmt.Println(err.Namespace())
			//fmt.Println(err.Field())
			//fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			//fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			//fmt.Println(err.Tag())
			//fmt.Println(err.ActualTag())
			//fmt.Println(err.Kind())
			//fmt.Println(err.Type())
			//fmt.Println(err.Value())
			//fmt.Println(err.Param())
			//fmt.Println()
		}
	}
	return nil
}

func (b BaseApiController) Hello(ctx *macaron.Context) {
	r := b.Failed(200, "")
	ctx.JSON(200, &r)
}
