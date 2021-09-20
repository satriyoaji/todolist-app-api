package repository

import (
	"context"
	"database/sql"
	"errors"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/domain"
	"time"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "insert into roles(name, description, created_at, updated_at) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, role.Name, role.Description, time.Now(), time.Now())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	role.Id = int(id)
	return role
}

func (repository *RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "update roles set name = ?, description = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, role.Name, role.Description, time.Now(), role.Id)
	helper.PanicIfError(err)

	return role
}

func (repository *RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, role domain.Role) {
	SQL := "delete from roles where id = ?"
	_, err := tx.ExecContext(ctx, SQL, role.Id)
	helper.PanicIfError(err)
}

func (repository *RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roleId int) (domain.Role, error) {
	SQL := "select id, name, description from roles where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roleId)
	helper.PanicIfError(err)
	defer rows.Close()

	role := domain.Role{}
	if rows.Next() {
		err := rows.Scan(&role.Id, &role.Name, &role.Description)
		helper.PanicIfError(err)
		return role, nil
	} else {
		return role, errors.New("role is not found")
	}
}

func (repository *RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Role {
	SQL := "select id, name, description from roles"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.Id, &role.Name, &role.Description)
		helper.PanicIfError(err)
		roles = append(roles, role)
	}
	return roles
}
