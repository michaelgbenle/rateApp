package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/helper"
)

//PingHandler is to check if server is up
func (u *HTTPHandler) PingHandler(c *gin.Context) {
	data := "server up"

	// healthcheck
	helper.Response(c, "pong", 200, data, nil)
}

//WelcomeHandler to welcome users
func (u *HTTPHandler) WelcomeHandler(c *gin.Context) {
	data := "Welcome to Exchange Rates WebApp Base Url by Gbenle Michael"

	// response
	helper.Response(c, "visit /register to sign up", 200, data, nil)
}
