package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/helper"
)

func (u *HTTPHandler) UsdNgnHandler(c *gin.Context) {
	data := "i'm ready"

	// healthcheck
	helper.Response(c, "pong", 200, data, nil)
}

func (u *HTTPHandler) NgnUsdHandler(c *gin.Context) {
	data := "i'm ready"

	// healthcheck
	helper.Response(c, "pong", 200, data, nil)
}
