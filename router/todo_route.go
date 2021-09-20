package router

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"satriyoaji/todolist-app-api/controller"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/repository"
	"satriyoaji/todolist-app-api/service"
)

func NewTodoRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) *httprouter.Router {

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService)

	router.GET("/api/todos", todoController.FindAll)
	router.GET("/api/todos/:todoId", todoController.FindById)
	router.POST("/api/todos", todoController.Create)
	router.PUT("/api/todos/:todoId", todoController.Update)
	router.DELETE("/api/todos/:todoId", todoController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
