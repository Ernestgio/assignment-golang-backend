package handler_test

import (
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

func TestHandlerGetTransactionWihtParams(t *testing.T) {
	t.Run("Should return 200 status code when request is valid", func(t *testing.T) {
		mockUsecase := new(mocks.TransactionUsecase)
		mockUsecase.On("GetWithParams", "created_at", "asc", "a", 10, 1).Return([]*entity.Transaction{}, nil)

		cfg := &server.RouterConfig{
			TransactionUsecase: mockUsecase,
		}

		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})

		req, _ := http.NewRequest("GET", "/transactions?sortBy=date&sort=asc&limit=10&s=a", nil)
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should return 200 status code when request is valid (case limit is default", func(t *testing.T) {
		mockUsecase := new(mocks.TransactionUsecase)
		mockUsecase.On("GetWithParams", "created_at", "asc", "a", 10, 1).Return([]*entity.Transaction{}, nil)

		cfg := &server.RouterConfig{
			TransactionUsecase: mockUsecase,
		}

		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})

		req, _ := http.NewRequest("GET", "/transactions?sortBy=date&sort=asc&s=a", nil)
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should return 500 status code when request is failed to processed by repo", func(t *testing.T) {
		mockUsecase := new(mocks.TransactionUsecase)
		mockUsecase.On("GetWithParams", "created_at", "asc", "a", 10, 1).Return(nil, sentinelerrors.ErrInternalServerError)

		cfg := &server.RouterConfig{
			TransactionUsecase: mockUsecase,
		}

		jwt, _ := hashutils.NewHashUtils().GenerateAccessToken(&entity.User{ID: 1, Wallet: &entity.Wallet{ID: 1}})

		req, _ := http.NewRequest("GET", "/transactions?sortBy=date&sort=asc&limit=10&s=a", nil)
		req.Header.Set("Authorization", "Bearer "+jwt.AccessToken)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
