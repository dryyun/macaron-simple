package form

type UserRegisterForm struct {
	Phone string `json:"phone" form:"phone" validate:"required,len=11"`
	//Password        string `json:"password" form:"password" validate:"required,min=6,eqfield=ConfirmPassword" `
	//ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,min=6"`
}

var UserRegisterFormErrorMsg = map[string]string{
	"Phone.required": "手机号必须填写",
	"Phone.len":      "手机号格式错误",
}
