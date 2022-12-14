package dto

type TopUpResponseDto struct {
	Amount              int    `json:"amount"`
	DestinationWalletId int    `json:"destination_wallet_id"`
	Description         string `json:"description"`
	TransactionStatus   string `json:"transaction_status"`
	SourceOfFundId      int    `json:"source_of_fund_id"`
}

type TopupRequestDto struct {
	Amount         int `json:"amount" binding:"required,gte=50000,lte=10000000"`
	SourceOfFundId int `json:"source_of_fund_id" binding:"required"`
}
