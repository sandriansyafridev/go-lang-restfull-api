package main

import (
	"golangapi/app"
	"golangapi/controller"
	"golangapi/repository"
	"golangapi/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	defer db.Close()

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	r := httprouter.New()
	r.GET("/api/categories", categoryController.FindAll)
	r.GET("/api/categories/:id", categoryController.FindByID)
	r.DELETE("/api/categories/:id", categoryController.Delete)

	http.ListenAndServe(":3000", r)

}
