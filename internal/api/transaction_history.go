package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/helper"
	"net/http"
)

func (u *HTTPHandler) TransactionsHandler(c *gin.Context) {
	user, err := u.GetUserFromContext(c)
	if err != nil {
		helper.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	//get transaction history and balances
	transactionHistory, err := u.Repository.GetTransactions(user)
	if err != nil {
		helper.Response(c, "error", 500, nil, []string{"internal server error"})
		return
	}
	// successful
	helper.Response(c, "transaction history and balances", 200, transactionHistory, nil)
}
