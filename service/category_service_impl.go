package service

import (
	"context"
	"golangapi/model/formatter"
	"golangapi/model/response"
	"golangapi/repository"
)

type categoryService struct {
	CategoryRepository repository.CategoryRepository
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
