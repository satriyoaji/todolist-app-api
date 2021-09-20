package helper

import (
	"satriyoaji/todolist-app-api/model/domain"
	masterRole "satriyoaji/todolist-app-api/model/web/role"
)

func ToRoleResponse(value domain.Role) masterRole.RoleResponse {
	return masterRole.RoleResponse{
		Id:          value.Id,
		Name:        value.Name,
		Description: value.Description,
		CreatedAt:   value.CreatedAt,
		UpdatedAt:   value.UpdatedAt,
	}
}

func ToRoleResponses(roles []domain.Role) []masterRole.RoleResponse {
	var roleResponses []masterRole.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}
