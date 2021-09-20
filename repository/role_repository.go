package repository

import (
	"context"
	"database/sql"
	"satriyoaji/todolist-app-api/model/domain"
)

type RoleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todo domain.Role) domain.Role
	Update(ctx context.Context, tx *sql.Tx, todo domain.Role) domain.Role
	Delete(ctx context.Context, tx *sql.Tx, todo domain.Role)
	FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Role, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Role
}
