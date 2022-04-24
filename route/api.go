package route

import (
	"golangapi/controller"

	"github.com/julienschmidt/httprouter"
)

func Initialize(categoryController controller.CategoryController) *httprouter.Router {

	r := httprouter.New()
	r.GET("/api/categories", categoryController.FindAll)
	r.POST("/api/categories", categoryController.Create)
	r.GET("/api/categories/:id", categoryController.FindByID)
	r.PUT("/api/categories/:id", categoryController.Update)
	r.DELETE("/api/categories/:id", categoryController.Delete)

	return r

}
