package service

import (
	"context"
	"satriyoaji/todolist-app-api/model/web/todo"
)

type TodoService interface {
	Create(ctx context.Context, request todo.TodoCreateRequest) todo.TodoResponse
	Update(ctx context.Context, request todo.TodoUpdateRequest) todo.TodoResponse
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) todo.TodoResponse
	FindAll(ctx context.Context) []todo.TodoResponse
}
