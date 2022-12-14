package appconstants

const (
	// Business rules
	NoRowsAffected = 0
	ZeroBalance    = 0
	NoSourceOfFund = 0
	NoSourceWallet = 0
	MinTopUpAmt    = 50000
	MaxTopUpAmt    = 10000000

	// JWT
	DaysTokenActive = 30
	HoursInADay     = 24

	//	Transaction Type Enum
	TopUpTransactionType = "Topup"

	// Topup Status Description
	TopupUncertain = ""
	TopupSuccess   = "Success"
	TopupFailed    = "Failed"

	// Description for Topup
	TopupDescription = "Top Up from %v"
)
