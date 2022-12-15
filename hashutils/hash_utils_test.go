package hashutils_test

import (
	"assignment-golang-backend/hashutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndSalt(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		expected error
	}{
		{
			name:     "Should return hashed password",
			password: "password",
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			hashUtil := hashutils.NewHashUtils()
			_, err := hashUtil.HashAndSalt(testCase.password)
			assert.Equal(t, testCase.expected, err)
		})
	}
}

func TestComparePassword(t *testing.T) {
	testCases := []struct {
		name       string
		hashedPwd  string
		inputPwd   string
		mockResult bool
		expected   bool
	}{
		{
			name:      "Should return false for invalid password",
			hashedPwd: "hashedPassword",
			inputPwd:  "password",
			expected:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := hashutils.NewHashUtils().ComparePassword(testCase.hashedPwd, testCase.inputPwd)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
