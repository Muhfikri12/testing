package shippingservice

import (
	"ecommers/repository"

	"go.uber.org/zap"
)

type ShippingService struct {
	RepoShipping repository.AllRepository
	Logger       *zap.Logger
}

func NewShippingService(RepoShipping repository.AllRepository, Logger *zap.Logger) ShippingService {
	return ShippingService{
		RepoShipping,
		Logger,
	}
}
