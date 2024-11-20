package promotionsservice

import (
	"ecommers/model"
	"ecommers/repository"
	"time"

	"go.uber.org/zap"
)

type PromotionsService struct {
	Repo repository.AllRepository
	Log  *zap.Logger
}

func NewPromotionsService(repo repository.AllRepository, log *zap.Logger) PromotionsService {
	return PromotionsService{
		Repo: repo,
		Log:  log,
	}
}

func (b *PromotionsService) GetallCampaign(status bool, days int) (*[]model.Promotions, error) {

	today := time.Now()
	end := time.Now().AddDate(0, 0, days)

	banners, err := b.Repo.PromotionRepo.GetAllCampaign(status)
	if err != nil {
		b.Log.Error("Error fetch banner Service: " + err.Error())
		return nil, err
	}

	bannersArr := []model.Promotions{}

	for _, banner := range *banners {
		if banner.StartDate.Before(end) && banner.EndDate.After(today) {
			bannersArr = append(bannersArr, banner)
		}
	}

	return &bannersArr, nil
}
