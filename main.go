package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"satriyoaji/todolist-app-api/app"
	"satriyoaji/todolist-app-api/controller"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/middleware"
	"satriyoaji/todolist-app-api/repository"
	"satriyoaji/todolist-app-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService)
	router := app.NewRouter(todoController)

	server := http.Server{
		Addr:    "localhost:"+app.GoDotEnvVariable("APP_PORT"),
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
