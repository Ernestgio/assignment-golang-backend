package handler

import (
	"assignment-golang-backend/sentinelerrors"
	"assignment-golang-backend/usecase"
	"assignment-golang-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase usecase.UserUsecase
}

type Config struct {
	UserUsecase usecase.UserUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		userUsecase: cfg.UserUsecase,
	}
}

func (h *Handler) HandleNotFound(c *gin.Context) {
	utils.ResponseWithError(c, http.StatusNotFound, sentinelerrors.ErrNotFound.Error())
}
