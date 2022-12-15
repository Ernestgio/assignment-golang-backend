package dto

type TransferDto struct {
	To          int    `json:"to" binding:"required"`
	Amount      int    `json:"amount" binding:"required,gte=1000,lte=50000000"`
	Description string `json:"description" binding:"max=35"`
}
