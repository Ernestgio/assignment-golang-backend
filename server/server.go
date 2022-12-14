package server

import (
	"assignment-golang-backend/db"
	"assignment-golang-backend/hashutils"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	userRepo := repository.NewUserRepository(&repository.UserUConfig{
		DB: db.Get(),
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepository: userRepo,
		HashUtil:       hashutils.NewHashUtils(),
	})

	sourceOfFundRepo := repository.NewSourceOfFundRepository(&repository.SourceOfFundRepositoryConfig{DB: db.Get()})
	walletRepo := repository.NewWalletRepository(&repository.WalletRepositoryConfig{DB: db.Get()})
	walletUsecase := usecase.NewWalletUsecase(&usecase.WalletUConfig{
		WalletRepository:       walletRepo,
		SourceOfFundRepository: sourceOfFundRepo,
	})

	return NewRouter(&RouterConfig{
		UserUsecase:   userUsecase,
		WalletUsecase: walletUsecase,
	})
}

func Init() {
	r := CreateRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
