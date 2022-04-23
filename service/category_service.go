package service

import (
	"context"
	"golangapi/model/request"
	"golangapi/model/response"
)

type CategoryService interface {
	FindAll(c context.Context) ([]response.CategoryResponse, error)
	FindByID(c context.Context, CategoryID int) (response.CategoryResponse, error)
	Delete(c context.Context, CategoryID int) error
	Create(c context.Context, categoryRequest request.CategoryCreateRequest) (response.CategoryResponse, error)
}
