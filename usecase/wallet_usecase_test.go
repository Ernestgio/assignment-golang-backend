package usecase_test

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	mocks "assignment-golang-backend/mocks/repository"
	"assignment-golang-backend/sentinelerrors"
	"assignment-golang-backend/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopup(t *testing.T) {
	testCases := []struct {
		name                 string
		walletId             int
		topUpAmt             int
		sourceOfFundId       int
		description          string
		formattedDescription string
		firstMockResult      *entity.SourceOfFund
		firstMockErr         error
		secondMockResult     *dto.TopUpResponseDto
		secondMockErr        error
		expectedResult       *dto.TopUpResponseDto
		expectedErr          error
	}{
		{
			name:                 "should return apropriate response and nil error when topup successful and get source_of_fund successful",
			walletId:             777777,
			topUpAmt:             500000,
			sourceOfFundId:       1,
			description:          "",
			formattedDescription: "Top Up from ",
			firstMockResult:      &entity.SourceOfFund{},
			firstMockErr:         nil,
			secondMockResult:     &dto.TopUpResponseDto{},
			secondMockErr:        nil,
			expectedResult:       &dto.TopUpResponseDto{Description: "", TransactionStatus: appconstants.TopupSuccess},
			expectedErr:          nil,
		},
		{
			name:                 "should return nil response and error when topup successful but get source_of_fund failed",
			walletId:             777777,
			topUpAmt:             500000,
			sourceOfFundId:       1,
			description:          "",
			formattedDescription: "Top Up from ",
			firstMockResult:      nil,
			firstMockErr:         sentinelerrors.ErrSourceOfFundIdNotExists,
			secondMockResult:     &dto.TopUpResponseDto{},
			secondMockErr:        nil,
			expectedResult:       nil,
			expectedErr:          sentinelerrors.ErrSourceOfFundIdNotExists,
		},
		{
			name:                 "should return nil response and error when topup failed but get source_of_fund succeed",
			walletId:             777777,
			topUpAmt:             500000,
			sourceOfFundId:       1,
			description:          "",
			formattedDescription: "Top Up from ",
			firstMockResult:      &entity.SourceOfFund{},
			firstMockErr:         nil,
			secondMockResult:     nil,
			secondMockErr:        errors.New("db failed"),
			expectedResult:       nil,
			expectedErr:          errors.New("db failed"),
		},
	}

	for _, testCase := range testCases {
		mockWalletRepo := mocks.NewWalletRepository(t)
		mockSOFRepo := mocks.NewSourceOfFundRepository(t)
		useCase := usecase.NewWalletUsecase(&usecase.WalletUConfig{
			WalletRepository:       mockWalletRepo,
			SourceOfFundRepository: mockSOFRepo,
		})
		mockSOFRepo.On("GetSourceOfFundById", testCase.sourceOfFundId).Return(testCase.firstMockResult, testCase.firstMockErr)

		if testCase.firstMockErr == nil {
			mockWalletRepo.On("Topup", testCase.walletId, testCase.topUpAmt, testCase.sourceOfFundId, testCase.formattedDescription).Return(testCase.secondMockResult, testCase.secondMockErr)
		}

		res, err := useCase.Topup(testCase.walletId, testCase.topUpAmt, testCase.sourceOfFundId)

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedResult, res)
	}
}

func TestTransfer(t *testing.T) {
	testCases := []struct {
		name                string
		sourceWalletId      int
		DestinationWalletId int
		transferDto         *dto.TransferDto
		firstMockResult     *entity.Wallet
		firstMockErr        error
		secondMockResult    *entity.Wallet
		secondMockErr       error
		thirdMockResult     *dto.TransferDto
		thirdMockErr        error
		expectedResult      *dto.TransferDto
		expectedErr         error
	}{
		{
			name:                "should return apropriate response and nil error when transfer successful and all repo successful",
			sourceWalletId:      0,
			DestinationWalletId: 0,
			transferDto:         &dto.TransferDto{},
			firstMockResult:     &entity.Wallet{},
			firstMockErr:        nil,
			secondMockResult:    &entity.Wallet{},
			secondMockErr:       nil,
			thirdMockResult:     &dto.TransferDto{},
			thirdMockErr:        nil,
			expectedResult:      &dto.TransferDto{},
			expectedErr:         nil,
		},
		{
			name:                "should return nil response and error when get source wallet failed",
			sourceWalletId:      0,
			DestinationWalletId: 0,
			transferDto:         &dto.TransferDto{},
			firstMockResult:     nil,
			firstMockErr:        sentinelerrors.ErrWalletNotExists,
			thirdMockResult:     &dto.TransferDto{},
			secondMockResult:    &entity.Wallet{},
			secondMockErr:       nil,
			thirdMockErr:        nil,
			expectedResult:      nil,
			expectedErr:         sentinelerrors.ErrWalletNotExists,
		},
		{
			name:                "should return nil response and error when balance less than transfer amount",
			sourceWalletId:      0,
			DestinationWalletId: 0,
			transferDto:         &dto.TransferDto{Amount: 1000000},
			firstMockResult:     &entity.Wallet{Amount: 20000},
			firstMockErr:        nil,
			secondMockResult:    &entity.Wallet{},
			secondMockErr:       nil,
			thirdMockResult:     nil,
			thirdMockErr:        nil,
			expectedResult:      nil,
			expectedErr:         sentinelerrors.ErrInsufficientBalance,
		},
		{
			name:                "should return nil response and error when transfer in repo failed",
			sourceWalletId:      0,
			DestinationWalletId: 0,
			transferDto:         &dto.TransferDto{},
			firstMockResult:     &entity.Wallet{},
			firstMockErr:        nil,
			secondMockResult:    &entity.Wallet{},
			secondMockErr:       nil,
			thirdMockResult:     nil,
			thirdMockErr:        errors.New("transfer error"),
			expectedResult:      nil,
			expectedErr:         errors.New("transfer error"),
		},
	}

	for _, testCase := range testCases {
		mockWalletRepo := mocks.NewWalletRepository(t)
		useCase := usecase.NewWalletUsecase(&usecase.WalletUConfig{
			WalletRepository: mockWalletRepo,
		})

		mockWalletRepo.On("GetWalletById", testCase.sourceWalletId).Return(testCase.firstMockResult, testCase.firstMockErr)

		if testCase.firstMockErr == nil {
			mockWalletRepo.On("GetWalletById", testCase.DestinationWalletId).Return(testCase.secondMockResult, testCase.secondMockErr)
		}

		if testCase.firstMockErr == nil && testCase.firstMockResult.Amount >= testCase.transferDto.Amount && testCase.secondMockErr == nil {
			mockWalletRepo.On("Transfer", testCase.sourceWalletId, testCase.transferDto).Return(testCase.thirdMockResult, testCase.thirdMockErr)
		}

		res, err := useCase.Transfer(testCase.sourceWalletId, testCase.transferDto)

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedResult, res)

	}
}

func TestGetWalletById(t *testing.T) {
	testCases := []struct {
		name           string
		id             int
		mockResult     *entity.Wallet
		mockErr        error
		expectedResult *entity.Wallet
		expectedErr    error
	}{
		{
			name:           "should return apropriate response and nil error when get wallet successful",
			id:             0,
			mockResult:     &entity.Wallet{},
			mockErr:        nil,
			expectedResult: &entity.Wallet{},
			expectedErr:    nil,
		},
		{
			name:           "should return nil response and error when get wallet failed",
			id:             0,
			mockResult:     nil,
			mockErr:        sentinelerrors.ErrWalletNotExists,
			expectedResult: nil,
			expectedErr:    sentinelerrors.ErrWalletNotExists,
		},
	}

	for _, testCase := range testCases {
		mockWalletRepo := mocks.NewWalletRepository(t)
		useCase := usecase.NewWalletUsecase(&usecase.WalletUConfig{
			WalletRepository: mockWalletRepo,
		})

		mockWalletRepo.On("GetWalletById", testCase.id).Return(testCase.mockResult, testCase.mockErr)

		res, err := useCase.GetWalletById(testCase.id)

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedResult, res)
	}
}
