package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model          `json:"-"`
	ID                  int           `gorm:"primary_key"  json:"id"`
	CreatedAt           time.Time     `json:"created_at"`
	Amount              int           `json:"amount"`
	SourceWalletId      *int          `json:"source_wallet_id,omitempty"`
	SourceWallet        *Wallet       `gorm:"foreignKey:SourceWalletId" json:"-"`
	DestinationWalletId int           `json:"destination_wallet_id"`
	DestinationWallet   *Wallet       `gorm:"foreignKey:DestinationWalletId" json:"-"`
	Description         string        `json:"description"`
	TransactionType     string        `json:"transaction_type"`
	SourceOfFundId      *int          `json:"source_of_fund_id"`
	SourceOfFund        *SourceOfFund `gorm:"foreignKey:SourceOfFundId" json:"source_of_fund,omitempty"`
}
