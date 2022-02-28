package main

import (
	"net/http"
	"rizalpurnama/golang-restful-api-pzn/app"
	"rizalpurnama/golang-restful-api-pzn/controller"
	"rizalpurnama/golang-restful-api-pzn/exception"
	"rizalpurnama/golang-restful-api-pzn/helper"
	"rizalpurnama/golang-restful-api-pzn/repository"
	"rizalpurnama/golang-restful-api-pzn/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Buat router
	router := httprouter.New()

	router.POST("/api/v1/categories", categoryController.Create)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.GET("/api/v1/categories", categoryController.FindAll)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	// Buat Error Handler
	router.PanicHandler = exception.ErrorHandler

	// Buat http server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
