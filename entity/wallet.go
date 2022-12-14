package entity

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model `json:"-"`
	ID         int `gorm:"primaryKey" json:"id"`
	Amount     int `json:"amount"`
	UserId     int `json:"user_id"`
}
