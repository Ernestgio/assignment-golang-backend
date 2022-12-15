package handler

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactionWithParams(c *gin.Context) {
	searchKey := c.DefaultQuery(appconstants.TransactionSearchKey, appconstants.TransactionDefaultSearch)
	sortColumnQuery := c.Query(appconstants.TransactionSortColumnKey)
	sortColumn := utils.QueryMapper(sortColumnQuery)

	sort := c.DefaultQuery(appconstants.TransactionSortOrderKey, appconstants.TransactionDefaultSortOrder)
	strLimit := c.Query(appconstants.TransactionLimitKey)

	var limit int
	if strLimit == "" {
		limit = appconstants.TransactionDefaultLimit
	} else {
		limit, _ = strconv.Atoi(strLimit)
	}

	transactions, err := h.transactionUsecase.GetWithParams(sortColumn, sort, searchKey, limit, c.GetInt(appconstants.WalletContextKey))
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, transactions)
}
