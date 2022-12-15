package utils

import "assignment-golang-backend/appconstants"

func QueryMapper(query string) (dbColumn string) {
	switch query {
	case appconstants.TransactionAmount:
		dbColumn = appconstants.TransactionAmountDbColumn
	case appconstants.TransactionDate:
		dbColumn = appconstants.TransactionDateDbColumn
	case appconstants.TransactionTo:
		dbColumn = appconstants.TransactionToDbColumn
	default:
		dbColumn = appconstants.TransactionDateDbColumn
	}
	return dbColumn
}
