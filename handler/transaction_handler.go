package handler

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactionWithParams(c *gin.Context) {
	searchKey := c.DefaultQuery(appconstants.TransactionSearchKey, appconstants.TransactionDefaultSearch)
	sortColumn := c.DefaultQuery(appconstants.TransactionSortColumnKey, appconstants.TransactionDefaultSortColumn)
	sortBy := c.DefaultQuery(appconstants.TransactionSortOrderKey, appconstants.TransactionDefaultSortOrder)

	transactions, err := h.transactionUsecase.GetWithParams(sortColumn, sortBy, searchKey, appconstants.TransactionDefaultLimit, c.GetInt("walletId"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, transactions)
}
