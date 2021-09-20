package service

import (
	"context"
	"satriyoaji/todolist-app-api/model/web/role"
)

type RoleService interface {
	Create(ctx context.Context, request role.RoleCreateRequest) role.RoleResponse
	Update(ctx context.Context, request role.RoleUpdateRequest) role.RoleResponse
	Delete(ctx context.Context, roleId int)
	FindById(ctx context.Context, roleId int) role.RoleResponse
	FindAll(ctx context.Context) []role.RoleResponse
}
