package service

import (
	"context"
	"satriyoaji/todolist-app-api/model/web/auth"
)

type authService interface {
	Login(ctx context.Context, request auth.AuthLoginRequest) auth.AuthResponse
	Logout(ctx context.Context)
}
