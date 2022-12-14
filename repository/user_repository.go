package repository

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	IsUserWithEmailExist(email string) bool
	Register(newUser *entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserById(id int) (*entity.User, error)
}
type userRepositoryImpl struct {
	db *gorm.DB
}

type UserUConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserUConfig) UserRepository {
	return &userRepositoryImpl{db: cfg.DB}
}

func (u *userRepositoryImpl) IsUserWithEmailExist(email string) bool {
	return u.db.Where("email = ?", email).Find(&entity.User{}).RowsAffected > appconstants.NoRowsAffected

}

func (u *userRepositoryImpl) Register(newUser *entity.User) (*entity.User, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newUser).Error; err != nil {
			return err
		}

		newWallet := &entity.Wallet{
			Amount: appconstants.ZeroBalance,
			UserId: newUser.ID,
		}

		if err := tx.Create(newWallet).Error; err != nil {
			return err
		}

		return nil
	})
	return newUser, err
}

func (u *userRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.Preload("Wallet").Where("email = ?", email).First(user).Error
	fmt.Println(user)
	return user, err
}

func (u *userRepositoryImpl) GetUserById(id int) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.Preload("Wallet").Where("id = ?", id).First(user).Error
	return user, err
}
