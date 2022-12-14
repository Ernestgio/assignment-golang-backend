package dto

import "github.com/golang-jwt/jwt"

type UserClaim struct {
	Id       int `json:"id"`
	WalletId int `json:"wallet_id"`
	jwt.StandardClaims
}
