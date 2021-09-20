package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/domain"
	"satriyoaji/todolist-app-api/model/web/todo"
	"satriyoaji/todolist-app-api/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository, DB *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *TodoServiceImpl) Create(ctx context.Context, request todo.TodoCreateRequest) todo.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := domain.Todo{
		UserId: request.UserId,
		Title:  request.Title,
	}

	todo = service.TodoRepository.Save(ctx, tx, todo)

	return helper.ToTodoResponse(todo)
}

func (service *TodoServiceImpl) Update(ctx context.Context, request todo.TodoUpdateRequest) todo.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todo.UserId = request.UserId
	todo.Title = request.Title

	todo = service.TodoRepository.Update(ctx, tx, todo)

	return helper.ToTodoResponse(todo)
}

func (service *TodoServiceImpl) Delete(ctx context.Context, todoId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodoRepository.Delete(ctx, tx, todo)
}

func (service *TodoServiceImpl) FindById(ctx context.Context, todoId int) todo.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTodoResponse(todo)
}

func (service *TodoServiceImpl) FindAll(ctx context.Context) []todo.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todos := service.TodoRepository.FindAll(ctx, tx)

	return helper.ToTodoResponses(todos)
}
