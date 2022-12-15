package server

import (
	"assignment-golang-backend/handler"
	"assignment-golang-backend/hashutils"
	"assignment-golang-backend/middleware"
	"assignment-golang-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase        usecase.UserUsecase
	WalletUsecase      usecase.WalletUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		UserUsecase:        cfg.UserUsecase,
		WalletUsecase:      cfg.WalletUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
	})

	m := middleware.NewMiddleware(&middleware.MiddlewareConfig{HashUtil: hashutils.NewHashUtils()})

	router.Static("/docs", "swagger-ui")
	router.POST("/register", m.LoginRegisterMiddleware(), h.Register)
	router.POST("/login", m.LoginRegisterMiddleware(), h.Login)

	userGroup := router.Group("/users")
	userGroup.Use(m.AuthMiddleware())
	{
		userGroup.GET("/", h.GetUserById)
	}

	transactionGroup := router.Group("/transactions")
	transactionGroup.Use(m.AuthMiddleware())
	{
		transactionGroup.GET("", h.GetTransactionWithParams)
		transactionGroup.POST("/topup", m.TopupMiddleware(), h.Topup)
		transactionGroup.POST("/transfer", m.TransferMiddleware(), h.Transfer)
	}

	router.NoRoute(h.HandleNotFound)
	return router
}
