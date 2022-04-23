package repository

import (
	"context"
	"database/sql"
	"golangapi/model/entity"
	"log"
)

type categoryRepository struct {
	DB *sql.DB
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

	return category, nil

}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
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
