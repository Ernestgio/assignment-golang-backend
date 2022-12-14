package usecase

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
)

type SourceOfFundUsecase interface{}

type sourceOfFundUsecaseImpl struct {
	sourceOfFundRepository repository.SourceOfFundRepository
}

type SourceOfFundUConfig struct {
	SourceOfFundRepository repository.SourceOfFundRepository
}

func NewSourceOfFundUsecase(cfg *SourceOfFundUConfig) SourceOfFundUsecase {
	return &sourceOfFundUsecaseImpl{sourceOfFundRepository: cfg.SourceOfFundRepository}
}

func (u *sourceOfFundUsecaseImpl) GetSourceOfFundById(id int) (*entity.SourceOfFund, error) {
	return u.sourceOfFundRepository.GetSourceOfFundById(id)
}
