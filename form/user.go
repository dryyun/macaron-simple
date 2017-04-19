package form

type UserRegisterForm struct {
	Phone           string `form:"phone" binding:"Required"`
	Password        string `form:"Password" binding:"Required"`
	ConfirmPassword string `form:"confirm_password" binding:"Required"`
}
