package entity

type SourceOfFund struct {
	ID   int    `gorm:"primaryKey" json:"-"`
	Name string `json:"name"`
}
