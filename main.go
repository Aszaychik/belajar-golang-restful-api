package main

import (
	"Aszaychik/belajar-restful-api/app"
	"Aszaychik/belajar-restful-api/controller"
	"Aszaychik/belajar-restful-api/helper"
	"Aszaychik/belajar-restful-api/middleware"
	"Aszaychik/belajar-restful-api/repository"
	"Aszaychik/belajar-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}