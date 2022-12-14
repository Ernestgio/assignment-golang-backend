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

type sourceOfFundRepositoryImpl struct {
	db *gorm.DB
}

type SourceOfFundRepositoryConfig struct {
	DB *gorm.DB
}

func NewSourceOfFundRepository(cfg *SourceOfFundRepositoryConfig) SourceOfFundRepository {
	return &sourceOfFundRepositoryImpl{db: cfg.DB}
}

func (s *sourceOfFundRepositoryImpl) GetSourceOfFundById(id int) (*entity.SourceOfFund, error) {
	sourceOfFund := &entity.SourceOfFund{}
	res := s.db.Where("id = ?", id).First(sourceOfFund)
	if res.RowsAffected == appconstants.NoRowsAffected {
		return nil, sentinelerrors.ErrSourceOfFundIdNotExists
	}
	return sourceOfFund, res.Error
}
