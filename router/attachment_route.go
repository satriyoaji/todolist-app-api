package router

import (
	"github.com/julienschmidt/httprouter"
	"satriyoaji/todolist-app-api/controller"
	"satriyoaji/todolist-app-api/exception"
)

func NewAttachmentRouter(router *httprouter.Router, attachmentController controller.AttachmentController) *httprouter.Router {

	router.GET("/api/attachments", attachmentController.FindAll)
	router.GET("/api/attachments/:attachmentId", attachmentController.FindById)
	router.POST("/api/attachments", attachmentController.Create)
	router.PUT("/api/attachments/:attachmentId", attachmentController.Update)
	router.DELETE("/api/attachments/:attachmentId", attachmentController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
