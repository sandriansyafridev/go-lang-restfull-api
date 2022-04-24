package main

import (
	"golangapi/app"
	"golangapi/controller"
	"golangapi/repository"
	"golangapi/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validate := validator.New()
	db := app.NewDB()
	defer db.Close()

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService, validate)

	r := httprouter.New()
	r.GET("/api/categories", categoryController.FindAll)
	r.POST("/api/categories", categoryController.Create)
	r.GET("/api/categories/:id", categoryController.FindByID)
	r.PUT("/api/categories/:id", categoryController.Update)
	r.DELETE("/api/categories/:id", categoryController.Delete)

	http.ListenAndServe(":3000", r)

}
