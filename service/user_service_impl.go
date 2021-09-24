package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/domain"
	"satriyoaji/todolist-app-api/model/web/auth"
	"satriyoaji/todolist-app-api/model/web/user"
	"satriyoaji/todolist-app-api/repository"
	"strconv"
	"time"
)

const SecretKey = "secret"

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, roleRepository repository.RoleRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		RoleRepository: roleRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request user.UserCreateRequest) user.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hashedPassword, _ := helper.HashPassword(request.Password)
	user := domain.User{
		Fullname:       request.Fullname,
		Email:          request.Email,
		Password:       hashedPassword,
		ForgotPassword: request.ForgotPassword,
		RoleId:         request.RoleId,
	}
	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request user.UserUpdateRequest) user.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Fullname = request.Fullname
	user.Email = request.Email
	hashedPassword, _ := helper.HashPassword(request.Password)
	user.Password = hashedPassword
	user.ForgotPassword = request.ForgotPassword
	user.RoleId = request.RoleId

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	userResponsesWithRole := service.assignRoleNameUser(ctx, tx, users)

	return userResponsesWithRole
}

func (service *UserServiceImpl) Login(ctx context.Context, request auth.AuthLoginRequest) auth.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{}
	user, err = service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if user.Id <= 0 {
		panic(exception.NewNotFoundError("User not found"))
	}

	//match hashed password
	match, err := helper.CheckPasswordHash(request.Password, user.Password)
	if !match {
		panic(exception.NewNotFoundError("invalid credential !"))
	}

	// assign role name
	users := []domain.User{user}
	userResponsesWithRole := service.assignRoleNameUser(ctx, tx, users)
	newUserResponse := userResponsesWithRole[0]

	//set JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(newUserResponse.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))
	helper.PanicIfError(err)

	return helper.ToAuthResponse(userResponsesWithRole[0], token)
}

func (service *UserServiceImpl) Logout(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//action delete jwt token
}

func (service *UserServiceImpl) assignRoleNameUser(ctx context.Context, tx *sql.Tx, users []domain.User) []user.UserResponse {
	listRoles := service.RoleRepository.FindAll(ctx, tx)
	roleMaps := map[int]string{}

	for _, val := range listRoles {
		roleMaps[val.Id] = val.Name
	}
	userResponses := make([]user.UserResponse, 0)

	for _, val := range users {
		userResponses = append(userResponses, user.UserResponse{
			Id:       val.Id,
			Fullname: val.Fullname,
			Email:    val.Email,
			RoleName: roleMaps[val.RoleId],
		})
	}

	return userResponses
}
