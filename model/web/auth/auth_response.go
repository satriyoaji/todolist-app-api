package auth

import "satriyoaji/todolist-app-api/model/web/user"

type AuthResponse struct {
	User  user.UserResponse `json:"user"`
	Token string            `json:"token"`
}
