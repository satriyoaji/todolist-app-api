package router

import (
	"github.com/julienschmidt/httprouter"
	"satriyoaji/todolist-app-api/controller"
	"satriyoaji/todolist-app-api/exception"
)

func NewRoleRouter(router *httprouter.Router, roleController controller.RoleController) *httprouter.Router {

	router.GET("/api/roles", roleController.FindAll)
	router.GET("/api/roles/:roleId", roleController.FindById)
	router.POST("/api/roles", roleController.Create)
	router.PUT("/api/roles/:roleId", roleController.Update)
	router.DELETE("/api/roles/:roleId", roleController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
