package entity

type SourceOfFund struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
