package router

import (
	"github.com/julienschmidt/httprouter"
	"satriyoaji/todolist-app-api/controller"
	"satriyoaji/todolist-app-api/exception"
)

func NewTodoRouter(router *httprouter.Router, todoController controller.TodoController) *httprouter.Router {

	router.GET("/api/todos", todoController.FindAll)
	router.GET("/api/todos/:todoId", todoController.FindById)
	router.POST("/api/todos", todoController.Create)
	router.PUT("/api/todos/:todoId", todoController.Update)
	router.DELETE("/api/todos/:todoId", todoController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
