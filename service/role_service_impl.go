package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/domain"
	"satriyoaji/todolist-app-api/model/web/role"
	"satriyoaji/todolist-app-api/repository"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoleService(roleRepository repository.RoleRepository, DB *sql.DB, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, request role.RoleCreateRequest) role.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := domain.Role{
		Name:        request.Name,
		Description: request.Description,
	}

	role = service.RoleRepository.Save(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Update(ctx context.Context, request role.RoleUpdateRequest) role.RoleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	role.Name = request.Name
	role.Description = request.Description

	role = service.RoleRepository.Update(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Delete(ctx context.Context, roleId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, roleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.RoleRepository.Delete(ctx, tx, role)
}

func (service *RoleServiceImpl) FindById(ctx context.Context, roleId int) role.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindById(ctx, tx, roleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) FindAll(ctx context.Context) []role.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles := service.RoleRepository.FindAll(ctx, tx)

	return helper.ToRoleResponses(roles)
}
