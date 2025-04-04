package repository

import (
	"database/sql"
	"log"

	"github.com/breadwithmeth/naliv_go/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	if db == nil {
		log.Fatal("Database connection is nil")
	}
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetCategories(business_id int) ([]*models.Category, error) {
	query := `
	SELECT c.category_id,
		   c.name,
		   c.parent_category,
		   c2.name as parent_category_name
	FROM categories c 
	LEFT JOIN categories c2 ON c2.category_id = c.parent_category 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.ParentCategory,
			&category.ParentCategoryName,
		); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}
