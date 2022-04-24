package controller

import (
	"golangapi/helper"
	"golangapi/model/request"
	"golangapi/model/response"
	"golangapi/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type categoryController struct {
	CategoryService service.CategoryService
	Validate        *validator.Validate
}

// Update implements CategoryController
func (categoryController *categoryController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CategoryID, _ := strconv.Atoi(params.ByName("id"))

	categoryRequest := request.CategoryUpdateRequest{}
	err := helper.ReadRequestBody(r, &categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail read  request body", err.Error(), response.EmptyObject{})
		helper.WriteRequestBody(w, responseError)
		return
	}

	err = categoryController.Validate.Struct(categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail update  request body", err.Error(), response.EmptyObject{})
		helper.WriteRequestBody(w, responseError)
		return
	}

	categoryRequest.ID = CategoryID

	categoryUpdated, err := categoryController.CategoryService.Update(r.Context(), categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail to update category", err.Error(), response.EmptyObject{})
		helper.WriteRequestBody(w, responseError)
		return
	}

	responseSuccess := response.BuildResponseSuccess("Category Updated", categoryUpdated)
	helper.WriteRequestBody(w, responseSuccess)
}

// Create implements CategoryController
func (categoryController *categoryController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryRequest := request.CategoryCreateRequest{}
	err := helper.ReadRequestBody(r, &categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail read  request body", err.Error(), response.EmptyObject{})
		helper.WriteRequestBody(w, responseError)
		return
	}

	err = categoryController.Validate.Struct(categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail read  request body", err.Error(), response.EmptyObject{})
		helper.WriteRequestBody(w, responseError)
		return
	}

	categoryResponse, err := categoryController.CategoryService.Create(r.Context(), categoryRequest)
	if err != nil {
		responseError := response.BuildResponseError("fail to create category", err.Error(), response.EmptyObject{})
		helper.WriteRequestBody(w, responseError)
		return
	}

	responseSuccess := response.BuildResponseSuccess("CategoryCreated", categoryResponse)
	helper.WriteRequestBody(w, responseSuccess)
}

// Delete implements CategoryController
func (categoryController *categoryController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	CategoryID, _ := strconv.Atoi(params.ByName("id"))
	data := map[string]interface{}{
		"deleted": false,
	}
	err := categoryController.CategoryService.Delete(r.Context(), CategoryID)
	if err != nil {
		responseError := response.BuildResponseError("Failed delete category", err.Error(), data)
		helper.WriteRequestBody(w, responseError)
		return
	}

	data["deleted"] = true
	responseSuccess := response.BuildResponseSuccess("Delete category", data)
	helper.WriteRequestBody(w, responseSuccess)
}

// FindByID implements CategoryController
func (categoryController *categoryController) FindByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CategoryID, _ := strconv.Atoi(params.ByName("id"))
	categoryResponse, err := categoryController.CategoryService.FindByID(r.Context(), CategoryID)

	if err != nil || categoryResponse.ID == 0 {
		responseError := response.BuildResponseError("Failed get category", err.Error(), response.EmptyObject{})
		helper.WriteRequestBody(w, responseError)
		return
	}

	responseSuccess := response.BuildResponseSuccess("Get category", categoryResponse)
	helper.WriteRequestBody(w, responseSuccess)
}

// FindAll implements CategoryController
func (categoryController *categoryController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoriesResponse, _ := categoryController.CategoryService.FindAll(r.Context())

	responseSuccess := response.BuildResponseSuccess("Get all categories", categoriesResponse)
	helper.WriteRequestBody(w, responseSuccess)
}

func NewCategoryController(categoryService service.CategoryService, validate *validator.Validate) CategoryController {
	return &categoryController{
		CategoryService: categoryService,
		Validate:        validate,
	}
}
