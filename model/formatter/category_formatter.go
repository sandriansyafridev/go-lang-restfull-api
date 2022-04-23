package formatter

import (
	"golangapi/model/entity"
	"golangapi/model/request"
	"golangapi/model/response"
)

func ToCategoryResponse(category entity.Category) (categoryResponse response.CategoryResponse) {

	categoryResponse.ID = category.ID
	categoryResponse.Name = category.Name

	return categoryResponse

}

func ToCategoriesResponse(categories []entity.Category) (categoriesResponse []response.CategoryResponse) {

	for _, category := range categories {
		categoryResponse := ToCategoryResponse(category)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse

}

func ToCategoryEntity(categoryRequest request.CategoryCreateRequest) (category entity.Category) {
	category.Name = categoryRequest.Name
	return category
}
