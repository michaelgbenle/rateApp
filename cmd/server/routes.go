package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/api"
	"github.com/michaelgbenle/rateApp/internal/middleware"
	"github.com/michaelgbenle/rateApp/internal/ports"
	"time"
)

//SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/api/v1")
	{
		r.GET("/ping", handler.PingHandler)
		r.POST("/register", handler.SignUpHandler)
		r.POST("/login", handler.LoginHandler)
	}

	// authorizeUser authorizes all authorized users handlers
	authorizeUser := r.Group("/user")
	authorizeUser.Use(middleware.AuthorizeUser(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		authorizeUser.PATCH("usdngn", handler.UsdNgnHandler)
		authorizeUser.PATCH("ngnusd", handler.NgnUsdHandler)
		authorizeUser.GET("transactions_history", handler.TransactionsHandler)
	}

	return router
}
