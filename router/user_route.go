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

func NewUserRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) *httprouter.Router {

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:userId", userController.FindById)
	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
