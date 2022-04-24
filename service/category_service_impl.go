package service

import (
	"context"
	"golangapi/model/entity"
	"golangapi/model/formatter"
	"golangapi/model/request"
	"golangapi/model/response"
	"golangapi/repository"
)

type categoryService struct {
	CategoryRepository repository.CategoryRepository
}

// Update implements CategoryService
func (categoryService *categoryService) Update(c context.Context, categoryRequest request.CategoryUpdateRequest) (response.CategoryResponse, error) {

	categoryResponse := response.CategoryResponse{}
	category, err := categoryService.CategoryRepository.FindByID(c, categoryRequest.ID)
	if err != nil {
		return categoryResponse, err
	}

	category.Name = categoryRequest.Name

	categoryUpdated, err := categoryService.CategoryRepository.Update(c, category)
	if err != nil {
		return categoryResponse, err
	}

	return formatter.ToCategoryResponse(categoryUpdated), nil

}

// Create implements CategoryService
func (categoryService *categoryService) Create(c context.Context, categoryRequest request.CategoryCreateRequest) (response.CategoryResponse, error) {
	categoryResponse := response.CategoryResponse{}
	category := entity.Category{}
	category.Name = categoryRequest.Name
	categoryCreated, err := categoryService.CategoryRepository.Create(c, category)
	if err != nil {
		return categoryResponse, err
	}

	return formatter.ToCategoryResponse(categoryCreated), nil

}

// Delete implements CategoryService
func (categoryService *categoryService) Delete(c context.Context, CategoryID int) error {

	category, err := categoryService.CategoryRepository.FindByID(c, CategoryID)
	if err != nil || category.ID == 0 {
		return err
	}

	err = categoryService.CategoryRepository.Delete(c, category)
	if err != nil {
		return err
	}

	return nil

}

// FindByID implements CategoryService
func (categoryService *categoryService) FindByID(c context.Context, CategoryID int) (response.CategoryResponse, error) {
	categoryResponse := response.CategoryResponse{}
	category, err := categoryService.CategoryRepository.FindByID(c, CategoryID)
	if err != nil {
		return categoryResponse, err
	}

	if category.ID == 0 {
		return categoryResponse, err
	}

	return formatter.ToCategoryResponse(category), nil
}

// FindAll implements CategoryService
func (categoryService *categoryService) FindAll(c context.Context) ([]response.CategoryResponse, error) {
	categoriesResponse := []response.CategoryResponse{}
	categories, err := categoryService.CategoryRepository.FindAll(c)
	if err != nil {
		return categoriesResponse, err
	}

	return formatter.ToCategoriesResponse(categories), nil
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		CategoryRepository: categoryRepository,
	}
}
