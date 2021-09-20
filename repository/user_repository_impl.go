package repository

import (
	"context"
	"database/sql"
	"errors"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/domain"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(fullname, email, password, forgot_password, role_id, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Fullname, user.Email, user.Password, user.ForgotPassword, user.RoleId, time.Now(), time.Now())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update users set fullname = ?, email = ?, password = ?, forgot_password = ?, role_id = ?,  updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Fullname, user.Email, user.Password, user.ForgotPassword, user.RoleId, time.Now(), user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from users where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := `select users.id, users.fullname, users.email, users.password, users.forgot_password, users.role_id, roles.name AS role
		from users 
		LEFT JOIN roles 
		ON users.role_id = roles.id
		where users.id = ?`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	role := domain.Role{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password, &user.ForgotPassword, &user.RoleId, &role.Name)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := `select users.id, users.fullname, users.email, users.password, users.forgot_password, users.role_id, roles.name AS role_name
		from users 
		LEFT JOIN roles 
		ON users.role_id = roles.id`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		role := domain.Role{}
		err := rows.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password, &user.ForgotPassword, &user.RoleId, &role.Name)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}
