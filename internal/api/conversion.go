package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/helper"
	"github.com/michaelgbenle/rateApp/internal/models"
	"net/http"
)

func (u *HTTPHandler) UsdNgnHandler(c *gin.Context) {
	user, err := u.GetUserFromContext(c)
	if err != nil {
		helper.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}
	var amount *models.Exchange
	err = c.ShouldBindJSON(&amount)
	if err != nil {
		helper.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}
	//check if user has enough balance
	if helper.InsufficientBalance(user.Balance["NGN"], amount.Amount) {
		helper.Response(c, "error", 400, nil, []string{"insufficient balance in ngn account"})
		return
	}
	//get exchange rates
	rates, err := helper.GetRates()
	if err != nil {
		helper.Response(c, "error", 500, nil, []string{"unable to fetch rates"})
		return
	}
	//buy usd at ask price(higher rate) usdcngn_
	rate := rates.Data.Rates.Usdcngn0.Rate
	value := helper.ConvertNgnToUsd(amount.Amount, rate)

	//update user balance and save transaction
	transaction, err := u.Repository.UpdateUserbalances(user, amount, value)
	if err != nil {
		helper.Response(c, "error", 500, nil, []string{"internal server error"})
		return
	}
	// success
	helper.Response(c, "success", 200, transaction, nil)
}

func (u *HTTPHandler) NgnUsdHandler(c *gin.Context) {
	user, err := u.GetUserFromContext(c)
	if err != nil {
		helper.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}
	var amount *models.Exchange
	err = c.ShouldBindJSON(&amount)
	if err != nil {
		helper.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}
	//check if user has enough balance
	if helper.InsufficientBalance(user.Balance["USD"], amount.Amount) {
		helper.Response(c, "error", 400, nil, []string{"insufficient balance in usd account"})
		return
	}
	//get exchange rates
	rates, err := helper.GetRates()
	if err != nil {
		helper.Response(c, "error", 500, nil, []string{"unable to fetch rates"})
		return
	}
	//sell usd at bid price(lower rate)
	rate := rates.Data.Rates.Usdcngn.Rate
	value := helper.ConvertNgnToUsd(amount.Amount, rate)

	//update user balance and save transaction
	transaction, err := u.Repository.UpdateUserbalances(user, amount, value)
	if err != nil {
		helper.Response(c, "error", 500, nil, []string{"internal server error"})
		return
	}
	// success
	helper.Response(c, "success", 200, transaction, nil)

}
