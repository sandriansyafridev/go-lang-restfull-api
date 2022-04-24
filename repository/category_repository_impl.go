package repository

import (
	"context"
	"database/sql"
	"errors"
	"golangapi/helper"
	"golangapi/model/entity"
)

type categoryRepository struct {
	DB *sql.DB
}

// Update implements CategoryRepository
func (categoryRepository *categoryRepository) Update(c context.Context, category entity.Category) (entity.Category, error) {

	querySQL := "UPDATE categories SET name = ? WHERE id = ?"
	stmt, err := categoryRepository.DB.Prepare(querySQL)
	helper.FatalIfError("Failed to prepare query SQL", err)

	_, err = stmt.ExecContext(c, category.Name, category.ID)
	helper.FatalIfError("Failed to ExecContext", err)

	return category, nil

}

// Create implements CategoryRepository
func (categoryRepository *categoryRepository) Create(c context.Context, category entity.Category) (entity.Category, error) {

	querySQL := "INSERT INTO categories(name) VALUES(?)"
	stmt, err := categoryRepository.DB.Prepare(querySQL)
	helper.FatalIfError("Failed to prepare query SQL", err)

	result, err := stmt.ExecContext(c, category.Name)
	helper.FatalIfError("Failed to ExecContext", err)

	id, _ := result.LastInsertId()
	CategoryID := int(id)
	category.ID = CategoryID

	return category, nil

}

// Delete implements CategoryRepository
func (categoryRepository *categoryRepository) Delete(c context.Context, category entity.Category) error {

	querySQL := "DELETE FROM categories WHERE id = ?"
	stmt, err := categoryRepository.DB.Prepare(querySQL)
	helper.FatalIfError("Failed to prepare query SQL", err)

	_, err = stmt.ExecContext(c, category.ID)
	helper.FatalIfError("Failed to ExecContext", err)

	return nil

}

// FindByID implements CategoryRepository
func (categoryRepository *categoryRepository) FindByID(c context.Context, CategoryID int) (entity.Category, error) {
	querySQL := "SELECT * FROM categories WHERE id = ?"
	stmt, err := categoryRepository.DB.Prepare(querySQL)
	helper.FatalIfError("Failed to prepare query SQL", err)

	rows, err := stmt.QueryContext(c, CategoryID)
	helper.FatalIfError("Failed to QueryContext", err)
	defer stmt.Close()

	defer rows.Close()
	category := entity.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		helper.FatalIfError("Failed scan categories", err)
	}

	if category.ID == 0 {
		return category, errors.New("no category")
	}

	return category, nil

}

// FindAll implements CategoryRepository
func (categoryRepository *categoryRepository) FindAll(c context.Context) ([]entity.Category, error) {
	querySQL := "SELECT * FROM categories"
	stmt, err := categoryRepository.DB.Prepare(querySQL)
	helper.FatalIfError("Failed to prepare query SQL", err)

	rows, err := stmt.QueryContext(c)
	helper.FatalIfError("Failed to execute query context", err)
	defer stmt.Close()

	defer rows.Close()
	categories := []entity.Category{}
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		helper.FatalIfError("Failed scan categories", err)
		categories = append(categories, category)
	}

	return categories, nil

}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}
