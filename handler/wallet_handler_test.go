package handler_test

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/hashutils"
	mocks "assignment-golang-backend/mocks/usecase"
	"assignment-golang-backend/sentinelerrors"
	"assignment-golang-backend/server"
	"assignment-golang-backend/testutils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerTopup(t *testing.T) {
	t.Run("Should return 200 status code when request is valid", func(t *testing.T) {
		mockWalletUsecase := new(mocks.WalletUsecase)
		mockWalletUsecase.On("Topup", 1, 100000, 1).Return(&dto.TopUpResponseDto{}, nil)

		cfg := &server.RouterConfig{
			WalletUsecase: mockWalletUsecase,
		}

		req, _ := http.NewRequest("POST", "/transactions/topup", testutils.MakeRequestBody(&dto.TopupRequestDto{
			Amount:         100000,
			SourceOfFundId: 1,
		}))
		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)
		assert.Equal(t, http.StatusOK, rec.Code)

	})
	t.Run("Should return 400 status code when request body is invalid", func(t *testing.T) {
		mockWalletUsecase := new(mocks.WalletUsecase)

		cfg := &server.RouterConfig{
			WalletUsecase: mockWalletUsecase,
		}

		req, _ := http.NewRequest("POST", "/transactions/topup", testutils.MakeRequestBody(&dto.TopupRequestDto{
			Amount: 100000}))
		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("Should return 500 status code when request is failed", func(t *testing.T) {
		mockWalletUsecase := new(mocks.WalletUsecase)
		mockWalletUsecase.On("Topup", 1, 100000, 1).Return(nil, sentinelerrors.ErrInternalServerError)

		cfg := &server.RouterConfig{
			WalletUsecase: mockWalletUsecase,
		}

		req, _ := http.NewRequest("POST", "/transactions/topup", testutils.MakeRequestBody(&dto.TopupRequestDto{
			Amount:         100000,
			SourceOfFundId: 1,
		}))
		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestHandlerTransfer(t *testing.T) {
	t.Run("Should return 200 status code when request is valid", func(t *testing.T) {
		mockWalletUsecase := new(mocks.WalletUsecase)
		mockWalletUsecase.On("Transfer", 1, &dto.TransferDto{
			Amount: 100000,
			To:     2,
		}).Return(&dto.TransferDto{}, nil)

		cfg := &server.RouterConfig{
			WalletUsecase: mockWalletUsecase,
		}

		req, _ := http.NewRequest("POST", "/transactions/transfer", testutils.MakeRequestBody(&dto.TransferDto{
			Amount: 100000,
			To:     2,
		}))
		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)
		assert.Equal(t, http.StatusOK, rec.Code)

	})

	t.Run("Should return 400 status code when request body is invalid", func(t *testing.T) {
		mockWalletUsecase := new(mocks.WalletUsecase)

		cfg := &server.RouterConfig{
			WalletUsecase: mockWalletUsecase,
		}

		req, _ := http.NewRequest("POST", "/transactions/transfer", testutils.MakeRequestBody(&dto.TransferDto{}))
		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Should return 200 status code when request is valid", func(t *testing.T) {
		mockWalletUsecase := new(mocks.WalletUsecase)
		mockWalletUsecase.On("Transfer", 1, &dto.TransferDto{
			Amount: 100000,
			To:     2,
		}).Return(nil, sentinelerrors.ErrInternalServerError)

		cfg := &server.RouterConfig{
			WalletUsecase: mockWalletUsecase,
		}

		req, _ := http.NewRequest("POST", "/transactions/transfer", testutils.MakeRequestBody(&dto.TransferDto{
			Amount: 100000,
			To:     2,
		}))
		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})
}
