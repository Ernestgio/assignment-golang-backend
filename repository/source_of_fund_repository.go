package repository

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/sentinelerrors"

	"gorm.io/gorm"
)

type SourceOfFundRepository interface {
	GetSourceOfFundById(id int) (*entity.SourceOfFund, error)
}

type SourceOfFundRepositoryImpl struct {
	db *gorm.DB
}

type SourceOfFundRepositoryConfig struct {
	DB *gorm.DB
}

func NewSourceOfFundRepository(cfg *SourceOfFundRepositoryConfig) SourceOfFundRepository {
	return &SourceOfFundRepositoryImpl{db: cfg.DB}
}

func (s *SourceOfFundRepositoryImpl) GetSourceOfFundById(id int) (*entity.SourceOfFund, error) {
	sourceOfFund := &entity.SourceOfFund{}
	res := s.db.Where("id = ?", id).First(sourceOfFund)
	if res.RowsAffected == appconstants.NoRowsAffected {
		return nil, sentinelerrors.ErrSourceOfFundIdNotExists
	}
	return sourceOfFund, res.Error
}
