package auth

type AuthLoginRequest struct {
	Email    string `validate:"required,email,min=1,max=100" json:"email" xml:"email" form:"email" `
	Password string `validate:"required,min=1,max=255" json:"password" xml:"password" form:"password"`
}
