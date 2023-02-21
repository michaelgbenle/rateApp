package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/helper"
)

//PingHandler is for testing the connections
func (u *HTTPHandler) PingHandler(c *gin.Context) {
	data := "server up"

	// healthcheck
	helper.Response(c, "pong", 200, data, nil)
}
