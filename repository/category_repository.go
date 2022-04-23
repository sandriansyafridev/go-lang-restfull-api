package repository

import (
	"context"
	"golangapi/model/entity"
)

type CategoryRepository interface {
	FindAll(c context.Context) ([]entity.Category, error)
	FindByID(c context.Context, CategoryID int) (entity.Category, error)
	Delete(c context.Context, cateogry entity.Category) error
}
