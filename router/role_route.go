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

func NewRoleRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) *httprouter.Router {

	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	router.GET("/api/roles", roleController.FindAll)
	router.GET("/api/roles/:roleId", roleController.FindById)
	router.POST("/api/roles", roleController.Create)
	router.PUT("/api/roles/:roleId", roleController.Update)
	router.DELETE("/api/roles/:roleId", roleController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
