package repository

import (
	"context"
	"database/sql"
	"errors"
	"golangapi/model/entity"
	"log"
)

type categoryRepository struct {
	DB *sql.DB
}

// Create implements CategoryRepository
func (categoryRepository *categoryRepository) Create(c context.Context, category entity.Category) (entity.Category, error) {

	querySQL := "INSERT INTO categories(name) VALUES(?)"
	stmt, err := categoryRepository.DB.Prepare(querySQL)
	if err != nil {
		log.Fatal("Failed to prepare query SQL:", err.Error())
	}

	result, err := stmt.ExecContext(c, category.Name)
	if err != nil {
		log.Fatal("Failed to ExecContext:", err.Error())
	}

	id, _ := result.LastInsertId()
	CategoryID := int(id)
	category.ID = CategoryID

	return category, nil

}

// Delete implements CategoryRepository
func (categoryRepository *categoryRepository) Delete(c context.Context, category entity.Category) error {

	querySQL := "DELETE FROM categories WHERE id = ?"
	stmt, err := categoryRepository.DB.Prepare(querySQL)
	if err != nil {
		log.Fatal("Failed to prepare query SQL:", err.Error())
	}

	_, err = stmt.ExecContext(c, category.ID)
	if err != nil {
		log.Fatal("Failed to ExecContext:", err.Error())
	}

	return nil

}

// FindByID implements CategoryRepository
func (categoryRepository *categoryRepository) FindByID(c context.Context, CategoryID int) (entity.Category, error) {
	querySQL := "SELECT * FROM categories WHERE id = ?"
	stmt, err := categoryRepository.DB.Prepare(querySQL)
	if err != nil {
		log.Fatal("Failed to prepare query SQL:", err.Error())
	}

	rows, err := stmt.QueryContext(c, CategoryID)
	if err != nil {
		log.Fatal("Failed to prepare query SQL:", err.Error())
	}
	defer stmt.Close()

	defer rows.Close()
	category := entity.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			log.Fatal("Failed scan categories:", err.Error())
		}
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
	if err != nil {
		log.Fatal("Failed to prepare query SQL:", err.Error())
	}

	rows, err := stmt.QueryContext(c)
	if err != nil {
		log.Fatal("Failed to execute query context:", err.Error())
	}

	defer stmt.Close()
	defer rows.Close()

	categories := []entity.Category{}
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			log.Fatal("Failed scan categories:", err.Error())
		}
		categories = append(categories, category)
	}

	return categories, nil

}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}
