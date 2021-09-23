package user

type UserResponse struct {
	Id        int    `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	RoleName  string `json:"role_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
