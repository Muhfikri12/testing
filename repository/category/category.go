package category

import (
	"database/sql"
	"ecommers/model"

	"go.uber.org/zap"
)

type CategoryRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewCategoryRepository(db *sql.DB, Log *zap.Logger) CategoryRepository {
	return CategoryRepository{
		DB:     db,
		Logger: Log,
	}
}

func (c *CategoryRepository) ShowAllCategory() (*[]model.Categories, error) {

	query := `SELECT id, name FROM categories WHERE deleted_at IS NULL`

	rows, err := c.DB.Query(query)
	if err != nil {
		c.Logger.Error("Error from repository: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	categories := []model.Categories{}

	for rows.Next() {
		category := model.Categories{}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			c.Logger.Error("Error from repository: " + err.Error())
			return nil, err
		}
		categories = append(categories, category)
	}

	return &categories, nil

}
