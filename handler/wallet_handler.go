package handler

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) Topup(c *gin.Context) {
	var request *dto.TopupRequestDto
	c.ShouldBindBodyWith(&request, binding.JSON)
	walletId := c.GetInt("walletId")
	fmt.Println("walletId: ", walletId)
	topupResponse, err := h.walletUsecase.Topup(walletId, request.Amount, request.SourceOfFundId)

	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, topupResponse)
}
