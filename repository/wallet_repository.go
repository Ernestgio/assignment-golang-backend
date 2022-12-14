package repository

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/sentinelerrors"

	"gorm.io/gorm"
)

type WalletRepository interface {
	GetWalletById(id int) (*entity.Wallet, error)
	Topup(walletId int, topUpAmt int, sourceOfFundId int, description string) (*dto.TopUpResponseDto, error)
}

type WalletRepositoryImpl struct {
	db *gorm.DB
}

type WalletRepositoryConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(cfg *WalletRepositoryConfig) WalletRepository {
	return &WalletRepositoryImpl{db: cfg.DB}
}

func (w *WalletRepositoryImpl) GetWalletById(id int) (*entity.Wallet, error) {
	wallet := &entity.Wallet{}
	res := w.db.Preload("User").Where("id = ?", id).First(wallet)
	if res.RowsAffected == appconstants.NoRowsAffected {
		return nil, sentinelerrors.ErrWalletNotExists
	}
	return wallet, res.Error
}

func (w *WalletRepositoryImpl) Topup(walletId int, topUpAmt int, sourceOfFundId int, description string) (*dto.TopUpResponseDto, error) {
	err := w.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Wallet{}).Where("id = ?", walletId).Update("amount", gorm.Expr("amount + ?", topUpAmt)).Error; err != nil {
			return err
		}
		if err := tx.Create(&entity.Transaction{
			Amount:              topUpAmt,
			DestinationWalletId: walletId,
			TransactionType:     appconstants.TopUpTransactionType,
			SourceOfFundId:      &sourceOfFundId,
			Description:         description,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	return &dto.TopUpResponseDto{
		Amount:              topUpAmt,
		SourceOfFundId:      sourceOfFundId,
		Description:         description,
		DestinationWalletId: walletId,
		TransactionStatus:   appconstants.TopupUncertain}, err
}
