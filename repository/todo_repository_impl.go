package repository

import (
	"context"
	"database/sql"
	"errors"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/domain"
	"time"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (repository *TodoRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "insert into todos(user_id, title, created_at, updated_at) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, todo.UserId, todo.Title, time.Now(), time.Now())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todo.Id = int(id)
	return todo
}

func (repository *TodoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "update todos set user_id = ?, title = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, todo.UserId, todo.Title, time.Now(), todo.Id)
	helper.PanicIfError(err)

	return todo
}

func (repository *TodoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo) {
	SQL := "delete from todos where id = ?"
	_, err := tx.ExecContext(ctx, SQL, todo.Id)
	helper.PanicIfError(err)
}

func (repository *TodoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error) {
	SQL := "select id, user_id, title from todos where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, todoId)
	helper.PanicIfError(err)
	defer rows.Close()

	todo := domain.Todo{}
	if rows.Next() {
		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("todo is not found")
	}
}

func (repository *TodoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo {
	SQL := "select id, user_id, title from todos"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var todos []domain.Todo
	for rows.Next() {
		todo := domain.Todo{}
		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}
	return todos
}
