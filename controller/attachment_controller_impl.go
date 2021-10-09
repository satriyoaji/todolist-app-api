package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/web"
	"satriyoaji/todolist-app-api/model/web/attachment"
	"satriyoaji/todolist-app-api/service"
	"strconv"
)

type AttachmentControllerImpl struct {
	AttachmentService service.AttachmentService
}

func NewAttachmentController(attachmentService service.AttachmentService) AttachmentController {
	return &AttachmentControllerImpl{
		AttachmentService: attachmentService,
	}
}

func (controller *AttachmentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	attachmentCreateRequest := attachment.AttachmentCreateRequest{}
	helper.ReadFromRequestBody(request, &attachmentCreateRequest)

	attachmentResponse := controller.AttachmentService.Create(request.Context(), attachmentCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   attachmentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AttachmentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	attachmentUpdateRequest := attachment.AttachmentUpdateRequest{}
	helper.ReadFromRequestBody(request, &attachmentUpdateRequest)

	attachmentId := params.ByName("attachmentId")
	id, err := strconv.Atoi(attachmentId)
	helper.PanicIfError(err)

	attachmentUpdateRequest.Id = id

	attachmentResponse := controller.AttachmentService.Update(request.Context(), attachmentUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   attachmentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AttachmentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	attachmentId := params.ByName("attachmentId")
	id, err := strconv.Atoi(attachmentId)
	helper.PanicIfError(err)

	controller.AttachmentService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AttachmentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	attachmentId := params.ByName("attachmentId")
	id, err := strconv.Atoi(attachmentId)
	helper.PanicIfError(err)

	attachmentResponse := controller.AttachmentService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   attachmentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AttachmentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	attachmentResponses := controller.AttachmentService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   attachmentResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
