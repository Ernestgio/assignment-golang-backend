package usecase

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	GetWithParams(sortBy string, sortDirection string, searchQuery string, limit int) ([]*entity.Transaction, error)
}

type transactionUseCaseImpl struct {
	transactionRepository repository.TransactionRepository
}

type TransactionUConfig struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(cfg *TransactionUConfig) TransactionUsecase {
	return &transactionUseCaseImpl{transactionRepository: cfg.TransactionRepository}
}

func (u *transactionUseCaseImpl) GetWithParams(sortBy string, sortDirection string, searchQuery string, limit int) ([]*entity.Transaction, error) {
	return u.transactionRepository.GetWithParams(sortBy, sortDirection, searchQuery, limit)
}
