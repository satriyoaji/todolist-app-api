package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/web"
	"satriyoaji/todolist-app-api/model/web/todo"
	"satriyoaji/todolist-app-api/service"
	"strconv"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (controller *TodoControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	todoCreateRequest := todo.TodoCreateRequest{}
	helper.ReadFromRequestBody(request, &todoCreateRequest)

	todoResponse := controller.TodoService.Create(request.Context(), todoCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TodoControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	todoUpdateRequest := todo.TodoUpdateRequest{}
	helper.ReadFromRequestBody(request, &todoUpdateRequest)

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoUpdateRequest.Id = id

	todoResponse := controller.TodoService.Update(request.Context(), todoUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TodoControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	controller.TodoService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TodoControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoResponse := controller.TodoService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TodoControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	todoResponses := controller.TodoService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
