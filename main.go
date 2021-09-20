package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"satriyoaji/todolist-app-api/app"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/middleware"
	"satriyoaji/todolist-app-api/router"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	mainRouter := router.NewRouter()

	// Todo's env
	mainRouter = router.NewTodoRouter(mainRouter, db, validate)

	// User's env
	mainRouter = router.NewUserRouter(mainRouter, db, validate)

	// Master Role's env
	mainRouter = router.NewRoleRouter(mainRouter, db, validate)

	server := http.Server{
		Addr:    "localhost:" + app.GoDotEnvVariable("APP_PORT"),
		Handler: middleware.NewAuthMiddleware(mainRouter),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
