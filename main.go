package main

import (
	"golangapi/app"
	"golangapi/controller"
	"golangapi/middleware"
	"golangapi/repository"
	"golangapi/route"
	"golangapi/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
	db       = app.NewDB()

	categoryRepository = repository.NewCategoryRepository(db)
	categoryService    = service.NewCategoryService(categoryRepository)
	categoryController = controller.NewCategoryController(categoryService, validate)
)

func main() {
	defer db.Close()
	r := route.Initialize(categoryController)
	http.ListenAndServe(":3000", middleware.NewAuthorizationMiddleware(r))
}
