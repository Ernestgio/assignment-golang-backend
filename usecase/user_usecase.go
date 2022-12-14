package usecase

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/hashutils"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/sentinelerrors"
)

type UserUsecase interface {
	Register(request *dto.UserDto) (*entity.User, error)
	Login(request *dto.UserDto) (*dto.LoginResponse, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
	hashUtil       hashutils.HashUtils
}

type UserUConfig struct {
	UserRepository repository.UserRepository
	HashUtil       hashutils.HashUtils
}

func NewUserUsecase(cfg *UserUConfig) UserUsecase {
	return &userUsecaseImpl{userRepository: cfg.UserRepository, hashUtil: cfg.HashUtil}
}

func (u *userUsecaseImpl) Register(request *dto.UserDto) (*entity.User, error) {
	if u.userRepository.IsUserWithEmailExist(request.Email) {
		return nil, sentinelerrors.ErrEmailAlreadyExists
	}
	hashedPwd, err := u.hashUtil.HashAndSalt(request.Password)

	if err != nil {
		return nil, err
	}

	newUser := &entity.User{}
	newUser.FromUserDto(request)
	newUser.Password = hashedPwd

	return u.userRepository.Register(newUser)
}

func (u *userUsecaseImpl) Login(request *dto.UserDto) (*dto.LoginResponse, error) {
	user, err := u.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, sentinelerrors.ErrEmailNotExists
	}

	if u.hashUtil.ComparePassword(user.Password, request.Password) {
		token, err := u.hashUtil.GenerateAccessToken(user)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
	return nil, sentinelerrors.ErrInvalidPassword
}
