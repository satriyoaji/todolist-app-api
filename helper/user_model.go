package helper

import (
	"satriyoaji/todolist-app-api/model/domain"
	"satriyoaji/todolist-app-api/model/web/user"
)

func ToUserResponse(value domain.User) user.UserResponse {
	return user.UserResponse{
		Id:             value.Id,
		Fullname:       value.Fullname,
		Email:          value.Email,
		Password:       value.Password,
		ForgotPassword: value.ForgotPassword,
		RoleId:         value.RoleId,
		CreatedAt:      value.CreatedAt,
		UpdatedAt:      value.UpdatedAt,
	}
}

func ToUserResponses(users []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
