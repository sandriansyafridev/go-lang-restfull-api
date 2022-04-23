package controller

import (
	"encoding/json"
	"golangapi/model/request"
	"golangapi/model/response"
	"golangapi/service"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type categoryController struct {
	CategoryService service.CategoryService
	Validate        *validator.Validate
}

// Create implements CategoryController
func (categoryController *categoryController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryRequest := request.CategoryCreateRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail read  request body", err.Error(), response.EmptyObject{})
		w.Header().Add("content-type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(responseError)
		return
	}

	err = categoryController.Validate.Struct(categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail read  request body", err.Error(), response.EmptyObject{})
		w.Header().Add("content-type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(responseError)
		return
	}

	categoryResponse, err := categoryController.CategoryService.Create(r.Context(), categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail to create category", err.Error(), response.EmptyObject{})
		w.Header().Add("content-type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(responseError)
		return
	}

	responseSuccess := response.BuildResponseSuccess("CategoryCreated", categoryResponse)
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(responseSuccess)

}

// Delete implements CategoryController
func (categoryController *categoryController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	CategoryID, _ := strconv.Atoi(params.ByName("id"))

	err := categoryController.CategoryService.Delete(r.Context(), CategoryID)
	if err != nil {
		responseError := response.BuildResponseError("Failed delete category", err.Error(), map[string]interface{}{
			"deleted": false,
		})
		w.Header().Add("content-type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(responseError)
		return
	}

	responseSuccess := response.BuildResponseSuccess("Delete category", map[string]interface{}{
		"deleted": true,
	})
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(responseSuccess)

}

// FindByID implements CategoryController
func (categoryController *categoryController) FindByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CategoryID, _ := strconv.Atoi(params.ByName("id"))
	categoryResponse, err := categoryController.CategoryService.FindByID(r.Context(), CategoryID)

	if err != nil || categoryResponse.ID == 0 {
		responseError := response.BuildResponseError("Failed get category", err.Error(), response.EmptyObject{})
		w.Header().Add("content-type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(responseError)
		return
	}

	responseSuccess := response.BuildResponseSuccess("Get category", categoryResponse)
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(responseSuccess)
}

// FindAll implements CategoryController
func (categoryController *categoryController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoriesResponse, err := categoryController.CategoryService.FindAll(r.Context())
	if err != nil {
		log.Fatal("NOT FOUND CATEGORIES")
	}
	responseSuccess := response.BuildResponseSuccess("Get all categories", categoriesResponse)
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(responseSuccess)

}

func NewCategoryController(categoryService service.CategoryService, validate *validator.Validate) CategoryController {
	return &categoryController{
		CategoryService: categoryService,
		Validate:        validate,
	}
}
