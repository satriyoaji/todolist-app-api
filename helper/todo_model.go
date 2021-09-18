package helper

import (
	"satriyoaji/todolist-app-api/model/domain"
	"satriyoaji/todolist-app-api/model/web/todo"
)

func ToTodoResponse(value domain.Todo) todo.TodoResponse {
	return todo.TodoResponse{
		Id: value.Id,
		UserId: value.UserId,
		Title: value.Title,
		CreatedAt: value.CreatedAt,
		UpdatedAt: value.UpdatedAt,
	}
}

func ToTodoResponses(todos []domain.Todo) []todo.TodoResponse {
	var todoResponses []todo.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, ToTodoResponse(todo))
	}
	return todoResponses
}
