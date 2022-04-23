package service

import (
	"context"
	"golangapi/model/formatter"
	"golangapi/model/response"
	"golangapi/repository"
	"log"
)

type categoryService struct {
	CategoryRepository repository.CategoryRepository
}

// Delete implements CategoryService
func (categoryService *categoryService) Delete(c context.Context, CategoryID int) error {

	category, err := categoryService.CategoryRepository.FindByID(c, CategoryID)
	log.Println(category)
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
