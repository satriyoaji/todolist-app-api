package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/web"
	"satriyoaji/todolist-app-api/model/web/user"
	"satriyoaji/todolist-app-api/service"
	"strconv"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	userCreateRequest := user.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	userUpdateRequest := user.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	_, err := helper.VerifyToken(request)
	if err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}
	_, err2 := helper.CheckRoleAdmin(request)
	if err2 != nil {
		panic(exception.UnauthorizedError{err2.Error()})
	}

	userResponses := controller.UserService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
