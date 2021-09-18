package app

import (
	"github.com/julienschmidt/httprouter"
	"satriyoaji/todolist-app-api/controller"
	"satriyoaji/todolist-app-api/exception"
)

func NewRouter(todoController controller.TodoController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/todos", todoController.FindAll)
	router.GET("/api/todos/:todoId", todoController.FindById)
	router.POST("/api/todos", todoController.Create)
	router.PUT("/api/todos/:todoId", todoController.Update)
	router.DELETE("/api/todos/:todoId", todoController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
