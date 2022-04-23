package formatter

import (
	"golangapi/model/entity"
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
