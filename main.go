package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"satriyoaji/todolist-app-api/app"
	"satriyoaji/todolist-app-api/controller"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/repository"
	"satriyoaji/todolist-app-api/router"
	"satriyoaji/todolist-app-api/service"
)

type Logger struct {
	handler *httprouter.Router
}

func (l *Logger) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

func main() {

	db := app.NewDB()
	validate := validator.New()
	mainRouter := router.NewRouter()

	// Todo's env
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService)

	// Master Role's env
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	// User's env
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, roleRepository, db, validate)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)

	// Master Attachment's env
	attachmentRepository := repository.NewAttachmentRepository()
	attachmentService := service.NewAttachmentService(attachmentRepository, db, validate)
	attachmentController := controller.NewAttachmentController(attachmentService)

	mainRouter = router.NewTodoRouter(mainRouter, todoController)
	mainRouter = router.NewUserRouter(mainRouter, userController)
	mainRouter = router.NewAuthRouter(mainRouter, authController)
	mainRouter = router.NewRoleRouter(mainRouter, roleController)
	mainRouter = router.NewAttachmentRouter(mainRouter, attachmentController)

	server := http.Server{
		Addr:    app.GoDotEnvVariable("APP_HOST_DEV") + ":" + app.GoDotEnvVariable("APP_PORT"),
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
