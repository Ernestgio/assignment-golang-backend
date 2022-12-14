package server

import (
	"assignment-golang-backend/handler"
	"assignment-golang-backend/middleware"
	"assignment-golang-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase usecase.UserUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		UserUsecase: cfg.UserUsecase,
	})

	m := middleware.NewMiddleware()

	router.POST("/register", m.LoginRegisterMiddleware(), h.Register)
	router.POST("/login", m.LoginRegisterMiddleware(), h.Login)

	router.NoRoute(h.HandleNotFound)
	return router
}
