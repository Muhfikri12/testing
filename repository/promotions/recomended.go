package promotions

import (
	"ecommers/model"
)

func (r *PromotionsRepository) GetAllRecomended() (*[]model.Promotions, error) {

	query := `SELECT id, title, image_url, subtitle, product_id FROM recomendeds 
		WHERE deleted_at IS NULL AND status = true`

	rows, err := r.DB.Query(query)
	if err != nil {
		r.Logger.Error("Error Query Data: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	recoments := []model.Promotions{}

	for rows.Next() {
		recoment := model.Promotions{}
		if err := rows.Scan(&recoment.ID, &recoment.Title, &recoment.ImageUrl, &recoment.Subtitle, &recoment.ProductId); err != nil {
			r.Logger.Error("Error Scan Data: " + err.Error())
			return nil, err
		}
		recoments = append(recoments, recoment)
	}

	return &recoments, nil
}
