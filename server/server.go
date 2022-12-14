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

	return NewRouter(&RouterConfig{
		UserUsecase: userUsecase,
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
