package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/web"
	"satriyoaji/todolist-app-api/model/web/auth"
	"satriyoaji/todolist-app-api/service"
)

type AuthControllerImpl struct {
	UserService service.UserService
}

func NewAuthController(authService service.UserService) AuthController {
	return &AuthControllerImpl{
		UserService: authService,
	}
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authLoginRequest := auth.AuthLoginRequest{}
	helper.ReadFromRequestBody(request, &authLoginRequest)

	authResponse := controller.UserService.Login(request.Context(), authLoginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}
	if errExpired := helper.MakeExpiredToken(request); errExpired != nil {
		panic(exception.UnauthorizedError{errExpired.Error()})
	}

	controller.UserService.Logout(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
