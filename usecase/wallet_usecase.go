package usecase

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/sentinelerrors"
	"fmt"
)

type WalletUsecase interface {
	GetWalletById(id int) (*entity.Wallet, error)
	Topup(walletId int, topUpAmt int, sourceOfFundId int) (*dto.TopUpResponseDto, error)
	Transfer(sourceWalletId int, transferDto *dto.TransferDto) (*dto.TransferDto, error)
}

type walletUsecaseImpl struct {
	walletRepository       repository.WalletRepository
	sourceOfFundRepository repository.SourceOfFundRepository
}

type WalletUConfig struct {
	WalletRepository       repository.WalletRepository
	SourceOfFundRepository repository.SourceOfFundRepository
}

func NewWalletUsecase(cfg *WalletUConfig) WalletUsecase {
	return &walletUsecaseImpl{walletRepository: cfg.WalletRepository, sourceOfFundRepository: cfg.SourceOfFundRepository}
}

func (u *walletUsecaseImpl) GetWalletById(id int) (*entity.Wallet, error) {
	return u.walletRepository.GetWalletById(id)
}

func (u *walletUsecaseImpl) Topup(walletId int, topUpAmt int, sourceOfFundId int) (*dto.TopUpResponseDto, error) {
	sourceOfFund, err := u.sourceOfFundRepository.GetSourceOfFundById(sourceOfFundId)
	if err != nil {
		return nil, err
	}

	description := fmt.Sprintf(appconstants.TopupDescription, sourceOfFund.Name)
	topupResult, err := u.walletRepository.Topup(walletId, topUpAmt, sourceOfFundId, description)
	if err != nil {
		return nil, err
	}
	topupResult.TransactionStatus = appconstants.TopupSuccess
	return topupResult, nil
}

func (u *walletUsecaseImpl) Transfer(sourceWalletId int, transferDto *dto.TransferDto) (*dto.TransferDto, error) {
	sourceWallet, err := u.walletRepository.GetWalletById(sourceWalletId)
	if err != nil {
		return nil, err
	}

	if sourceWallet.Amount < transferDto.Amount {
		return nil, sentinelerrors.ErrInsufficientBalance
	}

	destWallet, _ := u.walletRepository.GetWalletById(transferDto.To)
	if destWallet == nil {
		return nil, sentinelerrors.ErrDestinationWalletNotExists
	}

	transferResult, err := u.walletRepository.Transfer(sourceWalletId, transferDto)
	if err != nil {
		return nil, err
	}
	return transferResult, nil
}
