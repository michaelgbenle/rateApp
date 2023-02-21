package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/helper"
	"github.com/michaelgbenle/rateApp/internal/middleware"
	"github.com/michaelgbenle/rateApp/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

func (u *HTTPHandler) LoginHandler(c *gin.Context) {
	var loginRequest *models.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		helper.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}

	//check if email exists
	user, userErr := u.Repository.FindUserByEmail(loginRequest.Email)
	if userErr != nil {
		fmt.Println(userErr)
		helper.Response(c, "bad request", http.StatusBadRequest, nil, []string{"email does not exists"})
		return
	}

	//check if password is correct
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		helper.Response(c, "invalid Password", http.StatusBadRequest, nil, []string{"Bad Request"})
		return
	}

	// Generates access claims and refresh claims
	accessClaims, refreshClaims := middleware.GenerateClaims(user.Email)

	secret := os.Getenv("JWT_SECRET")
	accToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helper.Response(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helper.Response(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}
	c.Header("refresh_token", *refreshToken)
	c.Header("access_token", *accToken)

	helper.Response(c, "login successful", http.StatusOK, gin.H{
		"user":          user,
		"access_token":  *accToken,
		"refresh_token": *refreshToken,
	}, nil)
}
