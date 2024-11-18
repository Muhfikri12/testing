package repository

import (
	"database/sql"
	"ecommers/model"

	"go.uber.org/zap"
)

type SampelRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewSampelRepository(db *sql.DB, Log *zap.Logger) SampelRepository {
	return SampelRepository{
		DB:     db,
		Logger: Log,
	}
}

func (sampelRepository *SampelRepository) Create(book *model.Sampel) error {

	return nil
}

func (sampelRepository *SampelRepository) GetAll() ([]*model.Sampel, error) {
	return nil, nil
}

func (sampelRepository *SampelRepository) GetByID(id int) (*model.Sampel, error) {
	return nil, nil
}

func (sampelRepository *SampelRepository) Update(id int, book model.Sampel) error {
	return nil
}

func (sampelRepository *SampelRepository) Delete(id int) error {
	return nil
}
