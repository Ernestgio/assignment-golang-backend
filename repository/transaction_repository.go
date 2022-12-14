package repository

import (
	"assignment-golang-backend/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetWithParams(sortBy string, sortDirection string, searchQuery string, limit int) ([]*entity.Transaction, error)
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

type TransactionRepositoryConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(cfg *TransactionRepositoryConfig) TransactionRepository {
	return &transactionRepositoryImpl{db: cfg.DB}
}

func (s *transactionRepositoryImpl) GetWithParams(sortBy string, sortDirection string, searchQuery string, limit int) ([]*entity.Transaction, error) {
	transactions := []*entity.Transaction{}
	res := s.db.Where("description ILIKE ?", "%"+searchQuery+"%").Order(gorm.Expr("? ?", sortBy, sortDirection)).Limit(limit).Find(&transactions)
	return transactions, res.Error

}
