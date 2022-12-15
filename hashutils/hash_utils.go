package hashutils

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/config"
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type HashUtils interface {
	HashAndSalt(password string) (string, error)
	ComparePassword(hashedPwd string, inputPwd string) bool
	GenerateAccessToken(user *entity.User) (*dto.LoginResponse, error)
	ValidateToken(tokenString string) (*dto.UserClaim, error)
}

type hashUtilsImpl struct{}

func NewHashUtils() HashUtils {
	return &hashUtilsImpl{}
}

func (u *hashUtilsImpl) HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

func (u *hashUtilsImpl) ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}

func (u *hashUtilsImpl) GenerateAccessToken(user *entity.User) (*dto.LoginResponse, error) {
	claims := &dto.UserClaim{
		Id:       user.ID,
		WalletId: user.Wallet.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * appconstants.HoursInADay * appconstants.DaysTokenActive).Unix(),
			Issuer:    config.AppName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.SecretKey))

	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{AccessToken: tokenString}, nil
}

func (u *hashUtilsImpl) ValidateToken(tokenString string) (*dto.UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if ve, ok := err.(*jwt.ValidationError); ok {
		return nil, ve
	}
	claim := token.Claims.(*dto.UserClaim)

	return claim, nil
}
