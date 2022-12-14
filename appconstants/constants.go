package appconstants

const (
	// Business rules
	NoRowsAffected = 0
	ZeroBalance    = 0
	NoSourceOfFund = 0
	NoSourceWallet = 0

	// JWT
	DaysTokenActive = 30
	HoursInADay     = 24

	//	Transaction Type Enum
	TopUpTransactionType = "Topup"

	// Topup Status Description
	TopupUncertain = ""
	TopupSuccess   = "Success"

	// Description for Topup
	TopupDescription = "Top Up from %v"

	// List of Transactions Query Params Key
	TransactionSortOrderKey  = "sortBy"
	TransactionSortColumnKey = "sort"
	TransactionSearchKey     = "s"

	// List of Transactions Default Query Params Value
	TransactionDefaultLimit      = 10
	TransactionDefaultSortColumn = "created_at"
	TransactionDefaultSortOrder  = "desc"
	TransactionDefaultSearch     = ""
)
