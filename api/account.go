package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	db "simple-bank/db/sqlc"
)

type createAccountRequest struct {
	Owner        string `json:"owner" binding:"required"`
	CurrencyCode string `json:"currency_code" binding:"required,oneof=EUR GBP USD"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var request createAccountRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:        request.Owner,
		CurrencyCode: request.CurrencyCode,
		Balance:      decimal.NewFromInt(0),
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
