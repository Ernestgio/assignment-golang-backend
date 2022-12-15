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

	// Context Keys
	UserContextKey   = "userId"
	WalletContextKey = "walletId"

	//	Transaction Type Enum
	TopUpTransactionType = "Topup"

	// Topup Status Description
	TopupUncertain = ""
	TopupSuccess   = "Success"

	// Description for Topup
	TopupDescription = "Top Up from %v"

	// List of Transactions Query Params Key
	TransactionSortOrderKey  = "sort"
	TransactionSortColumnKey = "sortBy"
	TransactionSearchKey     = "s"
	TransactionLimitKey      = "limit"
	TransactionPageKey       = "page"

	// Allowed sort columns
	TransactionTo     = "to"
	TransactionDate   = "date"
	TransactionAmount = "amount"

	// Mapped sort columns (column name in database)
	TransactionToDbColumn     = "destination_wallet_id"
	TransactionDateDbColumn   = "created_at"
	TransactionAmountDbColumn = "amount"

	// List of Transactions Default Query Params Value
	TransactionDefaultLimit      = 10
	TransactionDefaultPage       = 1
	TransactionDefaultSortColumn = "created_at"
	TransactionDefaultSortOrder  = "desc"
	TransactionDefaultSearch     = ""
)
