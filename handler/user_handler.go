package handler

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/dto"
	"assignment-golang-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) Register(c *gin.Context) {
	var request dto.UserDto

	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.userUsecase.Register(&request)

	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {
	var request dto.UserDto

	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.userUsecase.Login(&request)

	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *Handler) GetUserById(c *gin.Context) {
	id := c.GetInt(appconstants.UserContextKey)
	user, err := h.userUsecase.GetUserById(id)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
