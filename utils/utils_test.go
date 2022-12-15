package utils_test

import (
	"assignment-golang-backend/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestMapper(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "Should return transaction amount db column",
			query:    "amount",
			expected: "amount",
		},
		{
			name:     "Should return transaction date db column",
			query:    "date",
			expected: "created_at",
		},
		{
			name:     "Should return transaction to db column",
			query:    "to",
			expected: "destination_wallet_id",
		},
		{
			name:     "Should return transaction date db column for invalid query",
			query:    "",
			expected: "created_at",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, utils.QueryMapper(testCase.query))
		})
	}
}
