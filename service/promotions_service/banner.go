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

func (b *PromotionsService) GetallBanners() (*[]model.Promotions, error) {

	today := time.Now()

	banners, err := b.Repo.PromotionRepo.GetAllBanner()
	if err != nil {
		b.Log.Error("Error fetch banner Service: " + err.Error())
		return nil, err
	}

	bannersArr := []model.Promotions{}

	for _, banner := range *banners {
		if !banner.StartDate.Before(today) { // Hanya banner dengan start_date hari ini atau lebih
			bannersArr = append(bannersArr, banner)
		}
	}

	return &bannersArr, nil
}

// func (b *PromotionsService) GetallBanners() (*[]model.Promotions, error) {
//
// 	banners, err := b.Repo.PromotionRepo.GetAllBanner()
// 	if err != nil {
// 		b.Log.Error("Error fatch banner Service: " + err.Error())
// 		return nil, err
// 	}
//
// 	return banners, nil
//
// }
