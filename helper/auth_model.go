package helper

import (
	"satriyoaji/todolist-app-api/model/web/auth"
	"satriyoaji/todolist-app-api/model/web/user"
)

func ToAuthResponse(value user.UserResponse, token string) auth.AuthResponse {

	authResponse := auth.AuthResponse{}
	authResponse.User = value
	authResponse.Token = token

	return auth.AuthResponse{
		User:  authResponse.User,
		Token: authResponse.Token,
	}
}
