package promotionsservice

import "ecommers/model"

func (rs *PromotionsService) GetAllRecomended() (*[]model.Promotions, error) {

	recoments, err := rs.Repo.PromotionRepo.GetAllRecomended()
	if err != nil {
		rs.Log.Error("Error From Service Recomendeds: " + err.Error())
		return nil, err
	}

	return recoments, nil
}
