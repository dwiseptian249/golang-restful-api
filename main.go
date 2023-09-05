package main

import (
	"net/http"
	"programmertio/golang-restful-api/app"
	"programmertio/golang-restful-api/controller"
	"programmertio/golang-restful-api/helper"
	"programmertio/golang-restful-api/middleware"
	"programmertio/golang-restful-api/repository"
	"programmertio/golang-restful-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
