package categoryservice

import (
	"ecommers/model"
	"ecommers/repository"

	"go.uber.org/zap"
)

type CategoriesService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewCategoriesService(repo repository.AllRepository, Log *zap.Logger) CategoriesService {
	return CategoriesService{
		Repo:   repo,
		Logger: Log,
	}
}

func (cs *CategoriesService) GetAllCategories() (*[]model.Categories, error) {

	categories, err := cs.Repo.CategoryRepo.ShowAllCategory()
	if err != nil {
		cs.Logger.Error("Error from Service: " + err.Error())
		return nil, err
	}

	return categories, nil

}
