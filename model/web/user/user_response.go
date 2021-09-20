package user

import "time"

type UserResponse struct {
	Id             int       `json:"id"`
	Fullname       string    `json:"fullname"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	ForgotPassword string    `json:"forgot_password"`
	RoleId         int       `json:"role_id"`
	RoleName       string    `json:"role_name"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
