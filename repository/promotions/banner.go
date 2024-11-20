package promotions

import (
	"database/sql"
	"ecommers/model"

	"go.uber.org/zap"
)

type PromotionsRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewPromotionsRepository(db *sql.DB, Log *zap.Logger) PromotionsRepository {
	return PromotionsRepository{
		DB:     db,
		Logger: Log,
	}
}

func (p *PromotionsRepository) GetAllBanner() (*[]model.Promotions, error) {

	query := `SELECT id, title, subtitle, image_url, path_url, start_date, end_date
		FROM promotions WHERE deleted_at IS NULL AND is_promo = false
		ORDER BY start_date ASC`

	rows, err := p.DB.Query(query)
	if err != nil {
		p.Logger.Error("Error from query banner repository: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	banners := []model.Promotions{}

	for rows.Next() {
		banner := model.Promotions{}
		if err := rows.Scan(&banner.ID, &banner.Title, &banner.Subtitle, &banner.ImageUrl, &banner.PathUrl, &banner.StartDate, &banner.EndDate); err != nil {
			p.Logger.Error("Error from looping banner repository: " + err.Error())
			return nil, err
		}

		banners = append(banners, banner)
	}

	return &banners, nil
}
