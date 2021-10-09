package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/web"
	"satriyoaji/todolist-app-api/model/web/role"
	"satriyoaji/todolist-app-api/service"
	"strconv"
)

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (controller *RoleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	roleCreateRequest := role.RoleCreateRequest{}
	helper.ReadFromRequestBody(request, &roleCreateRequest)

	roleResponse := controller.RoleService.Create(request.Context(), roleCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	roleUpdateRequest := role.RoleUpdateRequest{}
	helper.ReadFromRequestBody(request, &roleUpdateRequest)

	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	roleUpdateRequest.Id = id

	roleResponse := controller.RoleService.Update(request.Context(), roleUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	controller.RoleService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	helper.PanicIfError(err)

	roleResponse := controller.RoleService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RoleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	if _, err := helper.VerifyToken(request); err != nil {
		panic(exception.UnauthorizedError{err.Error()})
	}

	roleResponses := controller.RoleService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
