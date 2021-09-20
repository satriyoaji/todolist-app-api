package service

import (
	"context"
	"satriyoaji/todolist-app-api/model/web/user"
)

type UserService interface {
	Create(ctx context.Context, request user.UserCreateRequest) user.UserResponse
	Update(ctx context.Context, request user.UserUpdateRequest) user.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) user.UserResponse
	FindAll(ctx context.Context) []user.UserResponse
}
