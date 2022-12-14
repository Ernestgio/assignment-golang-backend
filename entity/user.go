package entity

import (
	"assignment-golang-backend/dto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         int     `gorm:"primary_key"  json:"id"`
	Email      string  `json:"name"`
	Password   string  `json:"-"`
	Wallet     *Wallet `gorm:"foreignKey:UserId" json:"wallet,omitempty"`
}

func (u *User) ToUserDto() *dto.UserDto {
	return &dto.UserDto{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *User) FromUserDto(dto *dto.UserDto) {
	u.Email = dto.Email
	u.Password = dto.Password
}
