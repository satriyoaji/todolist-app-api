package auth

type AuthRegisterRequest struct {
	Email                string `validate:"required,email,min=1,max=100" json:"email" xml:"email" form:"email" `
	Password             string `validate:"required,min=1,max=255" json:"password" xml:"password" form:"password"`
	PasswordConfirmation string `validate:"required,min=1,max=255" json:"password_confirmation" xml:"password_confirmation" form:"password_confirmation"`
}
