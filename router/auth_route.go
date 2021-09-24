package router

import (
	"github.com/julienschmidt/httprouter"
	"satriyoaji/todolist-app-api/controller"
	"satriyoaji/todolist-app-api/exception"
)

func NewAuthRouter(router *httprouter.Router, authController controller.AuthController) *httprouter.Router {

	router.POST("/api/login", authController.Login)
	router.POST("/api/logout", authController.Logout)

	router.PanicHandler = exception.ErrorHandler

	return router
}
