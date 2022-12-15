package usecase_test

import (
	"assignment-golang-backend/entity"
	mocks "assignment-golang-backend/mocks/repository"
	"assignment-golang-backend/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWithParams(t *testing.T) {
	testCases := []struct {
		name           string
		sortBy         string
		sortDirection  string
		searchQuery    string
		limit          int
		walletId       int
		mockResult     []*entity.Transaction
		mockErrResult  error
		expectedResult []*entity.Transaction
		expectedError  error
	}{
		{
			name:           "should return appropriate transactions and nil error when repo successfully get data",
			sortBy:         "column_name",
			sortDirection:  "desc",
			searchQuery:    "bank",
			limit:          10,
			walletId:       777777,
			mockResult:     []*entity.Transaction{},
			mockErrResult:  nil,
			expectedResult: []*entity.Transaction{},
			expectedError:  nil,
		},
	}

	for _, testCase := range testCases {
		mockRepo := mocks.NewTransactionRepository(t)
		useCase := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{
			TransactionRepository: mockRepo,
		})

		mockRepo.On("GetWithParams", testCase.sortBy, testCase.sortDirection, testCase.searchQuery, testCase.limit, testCase.walletId).Return(testCase.mockResult, testCase.mockErrResult)

		res, err := useCase.GetWithParams(testCase.sortBy, testCase.sortDirection, testCase.searchQuery, testCase.limit, testCase.walletId)

		assert.Equal(t, testCase.expectedResult, res)
		assert.Equal(t, testCase.expectedError, err)
	}
}
